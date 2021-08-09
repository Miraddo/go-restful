package main

import (
	"github.com/Miraddo/go-restful/c7/basicExample/helper"
	"log"
)

func main(){
	_, err := helper.InitDB()

	if err != nil{
		log.Println(err)
	}

	log.Println("Database tables are successfully initialized.")
}
