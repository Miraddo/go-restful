package main

import (
	"database/sql"
	"github.com/Miraddo/go-restful/c4/dbutils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB


// StationResource holds information about locations
type StationResource struct {
	ID int `json:"id"`
	Name string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func GetStation(ctx *gin.Context)  {
	var station StationResource

	id := ctx.Param("station_id")
	err := DB.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil{
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": station,
		})
	}
}

// CreateStation handles the POST

func CreateStation(c *gin.Context) {
	var station StationResource
	// Parse the body into our resource
	if err := c.BindJSON(&station); err == nil {
		// Format Time to Go time format
		statement, _ := DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)

		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}


// RemoveStation handles the removing of resource
func RemoveStation(c *gin.Context) {
	id := c.Param("station_id")
	statement, _ := DB.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}


func main()  {
	var err error
	// Connect to Database
	DB, err = sql.Open("sqlite3", "./c4/railAPI/railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	dbutils.Initialize(DB)

	r := gin.Default()
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	err = r.Run(":8888")
	if err != nil {
		log.Println(err)
		return 
	} // Default listen and serve on 0.0.0.0:8080
}