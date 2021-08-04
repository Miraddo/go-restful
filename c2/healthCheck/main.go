package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request)  {
	currentTime := time.Now()
	_, err := io.WriteString(w, currentTime.String())
	if err != nil{
		log.Fatal(err)
	}
}

func main()  {
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":8888", nil))
}