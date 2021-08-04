package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	// Mapping to methods is possible with HttpRouter
	router.ServeFiles("/static/*filepath", http.Dir("C:\\Users\\Miraddo\\go\\src\\go-restful\\c2\\fileServer\\static"))
	log.Fatal(http.ListenAndServe(":8888", router))
}
