package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// QueryArticleHandler Query Handler
func QueryArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Add this in your main program
	queryParams := r.URL.Query()

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Category is: %v\n", queryParams["category"][0])
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(w, "ID is: %v\n", queryParams["id"][0])

	if err != nil {
		panic(err)
	}
}

// ArticleHandler Path Handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(w, "ID is: %v\n", vars["id"])

	if err != nil {
		panic(err)
	}
}

func main() {
	routerMux := mux.NewRouter()
	// path method
	routerMux.HandleFunc("/articles1/{category}/{id:[0-9]+}", ArticleHandler)
	//query method
	routerMux.HandleFunc("/articles2", QueryArticleHandler)

	// path prefix
	routerMux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("C:\\Users\\Miraddo\\go\\src\\go-restful\\c2\\fileServer\\static"))))

	// Strict Slash
	routerMux.StrictSlash(true)
	routerMux.Path("/articles/").HandlerFunc(ArticleHandler)

	// Encoded Path
	routerMux.UseEncodedPath()
	routerMux.NewRoute().Path("/category/id")

	serv := &http.Server{
		Handler:      routerMux,
		Addr:         ":8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(serv.ListenAndServe())

}
