package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

// UUID is a custom multiplexer
type UUID struct{}

func (u *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/gen" {
		givenRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func givenRandomUUID(w http.ResponseWriter, r *http.Request) {

	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(w, fmt.Sprintf("%x", b))

	if err != nil {
		panic(err)
	}
}

func main() {
	mux := &UUID{}

	log.Fatal(http.ListenAndServe(":8888", mux))
}
