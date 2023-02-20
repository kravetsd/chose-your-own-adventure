package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arl/statsviz"
	"github.com/kravetsd/chose-your-own-adventure/cyoa"
)

func main() {
	//fl, err := os.Open("gopher.json")
	bts, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatal("Openning file:", err)
	}
	//defer fl.Close()

	// story, err := cyoa.JsonStory(fl)
	story, err := cyoa.JsonStoryDecode(bts)
	if err != nil {
		log.Fatal("Decoding json:", err)
	}

	sh := cyoa.NewStoryHandler(story, cyoa.WithTemplatePath("templates/story_new.html"), cyoa.WithUrlPath("/mysite"))
	http.Handle("/", sh)

	log.Println("Visit performance tool at http://localhost:8080/debug/statsviz/")
	statsviz.RegisterDefault()
	http.ListenAndServe(":8080", nil)

}
