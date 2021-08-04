package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware(originalHandler http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler
		originalHandler.ServeHTTP(w, r)

		fmt.Println("Executing middleware after response phase!")

	})
}

func handle(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Executing mainHandler...")
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
	return
}

func main()  {
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalHandler))

	log.Fatal(http.ListenAndServe(":8888", nil))

}