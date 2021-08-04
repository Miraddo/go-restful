package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Miraddo/go-restful/c2/urlShortener/pkg/store"
	"io"
	"log"
	"net/http"
	"strings"
)

var data []store.StoreUrls

// genKey generateKey
func genKey(u string) string {



	gen := base64.StdEncoding.EncodeToString([]byte(u))

	return gen
}

// getKey get mainUrl and return short key
func getKey(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "invalid_http_method")
		if err != nil {
			return 
		}
		return
	}

	var d int64

	err := r.ParseMultipartForm(d)

	if err != nil {
		return 
	}
	url := r.Form.Get("url")


	if url == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, x := range data{
		if x.Url == url{
			w.WriteHeader(http.StatusOK)
			_, err := io.WriteString(w, x.Key)

			if err != nil {
				return
			}
			return
		}
	}

	key := genKey(url)

	data = append(data, store.StoreUrls{Url: url, Key: key})
	fmt.Println(data)
	w.WriteHeader(http.StatusOK)
	_, err = io.WriteString(w, key)

	if err != nil {
		return
	}
	return
}

// getUrl get mainUrl and return short key
func getUrl(w http.ResponseWriter, r *http.Request)  {
	key := strings.TrimPrefix(r.URL.Path, "/api/v1/")

	for _, x := range data{
		if x.Key == key{
			w.WriteHeader(http.StatusOK)
			_, err := io.WriteString(w, x.Url)

			if err != nil {
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return
}



func main()  {

	http.HandleFunc("/api/v1/new", getKey)
	http.HandleFunc("/api/v1/", getUrl)

	log.Fatal(http.ListenAndServe(":8888", nil))

}
