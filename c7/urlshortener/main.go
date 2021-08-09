package main

import (
	"database/sql"
	"encoding/json"
	"github.com/Miraddo/go-restful/c7/urlshortener/helper"
	"github.com/Miraddo/go-restful/c7/urlshortener/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type DBClient struct {
	DB *sql.DB
}

type Record struct {
	ID int `json:"id"`
	URL string `json:"url"`
}


// GenerateShortURL adds URL to DB and gives back shortened string
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record

	postBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(postBody, &record)
	err = driver.DB.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.URL).Scan(&id)

	responseMap := map[string]string{"encoded_string": base62.ToBase62(id)}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		_, err = w.Write(response)
		if err != nil {
			return
		}
	}
}


// GetOriginalURL fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	// Get ID from base62 string
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.DB.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&url)
	// Handle response details
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		_, err = w.Write(response)

		if err != nil {
			return
		}
	}
}


func main()  {
	db, err := helper.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	dbclient := &DBClient{DB: db}
	if err != nil {
		log.Fatal(err)
	}

	defer func(cls *sql.DB) {
		err := cls.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Create a new router
	r := mux.NewRouter()

	// Attach an elegant path with handler
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")

	serv := &http.Server{
		Handler: r,
		Addr: ":8888",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(serv.ListenAndServe())
}
