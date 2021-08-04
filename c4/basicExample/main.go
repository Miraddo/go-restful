package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"time"
)

func pingTime(req *restful.Request, resp *restful.Response)  {
	// Write to the response
	_, err := io.WriteString(resp, fmt.Sprintf(
		"%s",
		time.Now(),
	))
	if err != nil {
		return 
	}
}

func main()  {
	// Create a web service
	webservice :=  new(restful.WebService)
	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// Add the service to application
	restful.Add(webservice)

	log.Fatal(http.ListenAndServe(":8888", nil))
}