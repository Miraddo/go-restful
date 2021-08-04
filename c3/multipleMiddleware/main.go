package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

// check content middleware
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// define a second middleware called
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setting cookie to every API response
		cookie := http.Cookie{Name: "ServerTimeUTC", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
		handler.ServeHTTP(w, r)
	})
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)

		if err != nil {
			log.Fatal(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(r.Body)

		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("201 - Created"))
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("405 - Method Not Allowed"))
		if err != nil {
			return
		}
	}
}

func main() {
	originalHandler := http.HandlerFunc(postHandler)

	http.Handle("/city", filterContentType(setServerTimeCookie(originalHandler)))

	log.Fatal(http.ListenAndServe(":8888", nil))
}
