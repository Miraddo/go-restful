package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type DB struct {
	collection *mongo.Collection
}

type Movie struct {
	ID interface{} `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Year string `json:"year" bson:"year"`
	Directors []string `json:"directors" bson:"directors"`
	Writers []string `json:"writers" bson:"writers"`
	BoxOffice BoxOffice `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross uint64 `json:"gross" bson:"gross"`
}

func (db *DB) GetMovie (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	var movie Movie

	objectID, _ := primitive.ObjectIDFromHex(vars["id"])

	filter := bson.M{"_id": objectID}

	err := db.collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(response)
		if err != nil {
			return
		}
	}
}




func (db *DB) PostMovie (w http.ResponseWriter, r *http.Request)  {

	var movie Movie
	postBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(postBody, &movie)
	if err != nil {
		return 
	}

	result, err := db.collection.InsertOne(context.TODO(), movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(response)
		if err != nil {
			return
		}
	}
}


// UpdateMovie modifies the data of given resource
func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie
	putBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(putBody, &movie)
	if err != nil {
		return
	}
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": &movie}
	_, err = db.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}

	} else {
		w.WriteHeader(http.StatusOK)
	}
}


// DeleteMovie removes the data from the db
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": objectID}
	_, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}

	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func main()  {
	aut := options.Credential{
		Username:      "root",
		Password:      "root",
	}
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017").SetAuth(aut)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal("Disconnect Error", err)
		}
	}(client, context.TODO())

	collection := client.Database("appDB").Collection("movies")
	db := &DB{collection: collection}

	r := mux.NewRouter()
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods("GET")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.UpdateMovie).Methods("PUT")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/v1/movies", db.PostMovie).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}