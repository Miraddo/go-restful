package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main(){
	r := gin.Default()

	r.GET("/pingTime", func(context *gin.Context) {
		// JSON serializer is available on gin context
		context.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	err := r.Run(":8888")
	if err != nil{
		log.Fatalln(err)
		return
	}
}