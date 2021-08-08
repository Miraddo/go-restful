package main

import (
	"database/sql"
	"encoding/json"
	"github.com/Miraddo/go-restful/c4/dbutils"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// POST 	/v1/train 		(details as JSON body) 		Create Train
// POST 	/v1/station 	(details as JSON body) 		Create Station
// GET 		/v1/train/id 								Read Train
// GET 		/v1/station/id 								Read Station
// POST 	/v1/schedule 	(source and destination) 	Create Route

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}



// Register adds paths and routes to a new service instance
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		err := response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
		if err != nil {
			return
		}
	} else {
		err := response.WriteEntity(t)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// POST http://localhost:8000/v1/trains
func (t *TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)

	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)

	// Error handling is obvious here. So omitting...
	statement, _ := DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values (?, ?)")

	result, err := statement.Exec(b.DriverName, b.OperatingStatus)

	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		err := response.WriteHeaderAndEntity(http.StatusCreated, b)
		if err != nil {
			return
		}
	}else{
		response.AddHeader("Content-type", "text/plain")
		err := response.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			return 
		}
	}


}

// Delete http://localhost:8000/v1/trains/1
func (t *TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ :=  DB.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)

	if err == nil{
		response.WriteHeader(http.StatusOK)
	}else{
		response.AddHeader("Content-Type", "text/plain")
		err := response.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			return 
		}
	}
}


func main() {
	var err error
	// Connect to Database
	DB, err = sql.Open("sqlite3", "./c4/railAPI/railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	// Create tables
	dbutils.Initialize(DB)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}

	t.Register(wsContainer)
	log.Printf("start listening on localhost:8888")
	server := &http.Server{Addr: ":8888", Handler: wsContainer}

	log.Fatal(server.ListenAndServe())
}
