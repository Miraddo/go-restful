package main

import (
	"encoding/json"
	"fmt"
	"github.com/Miraddo/go-restful/c1/mirrorFinder/pkg/mirrors"
	"log"
	"net/http"
	"time"
)

//  response url | latency
type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

// findFastest response to return fastest mirror
func findFastest(list []string) response {

	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range list {
		mirrorUrl := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorUrl + "/README")
			latency := time.Now().Sub(start)
			if err == nil {
				urlChan <- mirrorUrl
				latencyChan <- latency
			}

		}()
	}

	return response{<-urlChan, <-latencyChan}
}

func main() {

	// create handle request
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(respJson)
		if err != nil {
			log.Fatalln("Response JSON not Worked, Error : ", err)
		}
	})

	// run web server
	port := ":8888"

	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server on port %s \n", port)
	log.Fatal(server.ListenAndServe())
}

