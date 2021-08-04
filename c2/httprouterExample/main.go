package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	io "io"
	"log"
	"net/http"
	"os/exec"
)

func getCommandOutput(c string, args ...string) string {
	out, _ := exec.Command(c, args...).Output()

	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// I used windows path but for unix based system we can use this path /usr/local/go/bin/go
	res := getCommandOutput("C:\\Program Files\\Go\\bin\\go.exe", "version")
	_, err := io.WriteString(w, res)
	if err != nil {
		panic(err)
	}
	return
}

func getFileContent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := fmt.Fprintf(w, getCommandOutput("cat", p.ByName("name")))
	if err != nil {
		return
	}
}

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8888", router))
}
