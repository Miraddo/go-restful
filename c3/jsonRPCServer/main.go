package main

import (
	jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Args holds arguments passed to JSON-RPC service
type Args struct {
	ID string
}

// Book struct hoolds book JSON structure
type Book struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	// Read JSON file and load data
	absPath, _ := filepath.Abs("C:\\Users\\Miraddo\\go\\src\\go-restful\\c3\\jsonRPCServer\\books.json")

	raw, readerr := ioutil.ReadFile(absPath)
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}

	// Unmarshal JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over each book to find the given book

	for _, book := range books{
		if book.ID == args.ID{
			*reply = book
			break
		}
	}
	return nil
}

func main()  {
	// Create a new RPC server
	s := rpc.NewServer()

	// Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	err := s.RegisterService(new(JSONServer), "")
	if err != nil {
		return
	}
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	err = http.ListenAndServe(":8889", r)
	if err != nil {
		return
	}
}