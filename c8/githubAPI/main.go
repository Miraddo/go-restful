package main

import (
	"github.com/levigross/grequests"
	"log"
)

var GITHUB_TOKEN = "KEY"
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

type Repo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Forks int `json:"forks"`
	Private bool `json:"private"`
}

func getStats(url string) *grequests.Response{
	resp, err := grequests.Get(url, requestOptions)
	// You can modify the request by passing an optional RequestOptions struct
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}


func main() {
	var repos []Repo
	var repoUrl = "https://api.github.com/users/Miraddo/repos"
	resp := getStats(repoUrl)
	err := resp.JSON(&repos)
	if err != nil {
		return
	}
	log.Println(repos)
}