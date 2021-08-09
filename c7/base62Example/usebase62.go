package main

import (
	"github.com/Miraddo/go-restful/c7/base62Example/base62"
	"log"
)

func main() {
	x := 999999999999
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	normalNumber := base62.ToBase10(base62String)
	log.Println(normalNumber)

	tbe64 := base62.TBE62(x)
	log.Println(tbe64)

	tbd64 := base62.TBD62(tbe64)
	log.Println(tbd64)
}



