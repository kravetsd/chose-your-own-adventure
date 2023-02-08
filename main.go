package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kravetsd/chose-your-own-adventure/cyoa"
)

func main() {
	fl, err := os.Open("gopher.json")
	if err != nil {
		log.Fatal("Openning file:", err)
	}
	defer fl.Close()

	story, err := cyoa.JsonStory(fl)
	if err != nil {
		log.Fatal("Decoding json:", err)
	}

	sh := cyoa.NewStoryHandler(story)
	http.Handle("/", sh)

	http.ListenAndServe(":8080", nil)

}
