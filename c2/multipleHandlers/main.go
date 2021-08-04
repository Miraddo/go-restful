package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, rand.Float64())
		if err != nil{
			panic(err)
		}
	})

	newMux.HandleFunc("/randomInt", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, rand.Intn(100))
		if err != nil{
			panic(err)
		}
	})

	log.Fatalln(http.ListenAndServe(":8888", newMux))
}
