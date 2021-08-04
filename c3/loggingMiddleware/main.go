package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type city struct {
	Name string
	Area uint64
}


func postHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return 
	}
	log.Println("Finished processing request")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", postHandler)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8888", loggedRouter))
}
