package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Movie holds a movie data
type Movie struct {
	Name string `bson:"name"`
	Year int `bson:"year"`
	Directors []string `bson:"directors"`
	Writers []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

// BoxOffice is nested in Movie
type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross uint64 `bson:"gross"`
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

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Ping", err)
	}

	fmt.Println("Connected to MongoDB successfully")

	collection := client.Database("appDB").Collection("movies")

	// Create a movie
	darkNight := Movie{
		Name: "The Dark Knight",
		Year: 2008,
		Directors: []string{"Christopher Nolan"},
		Writers: []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross: 533316061,
		},
	}
	// Insert a document into MongoDB
	_, err = collection.InsertOne(context.TODO(), darkNight)
	if err != nil {
		log.Fatal("InsertOne", err)
	}

	queryResult := &Movie{}
	// bson.M is used for building map for filter query
	filter := bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}
	result := collection.FindOne(context.TODO(), filter)
	err = result.Decode(queryResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie:", queryResult)

	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("Disconnected from MongoDB")



}