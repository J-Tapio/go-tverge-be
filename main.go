package main

import "log"

type mainStory struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	URL    string `json:"URL"`
	Image  string `json:"Image"`
}

type feedStory struct {
	Title    string `json:"Title"`
	Author   string `json:"Author"`
	URL      string `json:"URL"`
	Date     string `json:"Date"`
	Image    string `json:"Image"`
	Comments string `json:"Comments"`
}

type featuredStory struct {
	Title     string `json:"Title"`
	Author    string `json:"Author"`
	URL       string `json:"URL"`
	Date      string `json:"Date"`
	Image     string `json:"Image"`
	PullQuote string `json:"PullQuote"`
}

type asideVideo struct {
	Title string `json:"Title"`
	URL   string `json:"URL"`
	Image string `json:"Image"`
}

type data struct {
	Image    string           `json:"BackgroundImg"`
	Quote    string           `json:"Quote"`
	Videos   []*asideVideo    `json:"Videos"`
	Main     []*mainStory     `json:"MainNews"`
	Feed     []*feedStory     `json:"FeedNews"`
	Featured []*featuredStory `json:"FeedFeatured"`
}

var coverImage string
var quote string
var feedAsideVideos []*asideVideo
var mainStoryData []*mainStory
var feedStoryData []*feedStory
var featuredStoryData []*featuredStory
var currentNews data

func main() {
	log.Println("Starting the server")
	go startScraper()
	startServer()
	log.Println("Server has stopped running")
}
