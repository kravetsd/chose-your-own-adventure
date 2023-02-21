package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arl/statsviz"
	"github.com/kravetsd/chose-your-own-adventure/cyoa"
)

func main() {

	//fl, err := os.Open("gopher.json")
	//To read from file and use io.Reader you can use cyoa.JsonStory()
	//defer fl.Close()
	// story, err := cyoa.JsonStory(fl)

	//JsonStoryDecode uses decoding from bytes
	// Please comment the next few lines and uncomment  lines above to swith to io.Read approach
	bts, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatal("Openning file:", err)
	}

	story, err := cyoa.JsonStoryDecode(bts)
	if err != nil {
		log.Fatal("Decoding json:", err)
	}

	ch := make(chan string)
	go func() {
		sh := cyoa.NewStoryHandler(story, cyoa.WithTemplatePath("templates/story_new.html"), cyoa.WithUrlPath(customFuncPath))
		statsviz.RegisterDefault()
		http.Handle("/mysite/", sh)
		log.Println("\n------------\nVisit performance tool at http://localhost:8080/debug/statsviz/\nYou also can read this story via web interface at http://localhost:8080/mysite/\n------------")
		http.ListenAndServe(":8080", nil)

	}()
	go func() {
		fmt.Println("Please enter your name")
		var name string
		fmt.Scan(&name)
		fmt.Printf("Hello, %v! ", name)
		ch <- name
	}()
	if <-ch == "done" {
		fmt.Println("Bye-bye!")
	}
}

// This is a custom function for handling custom url path for serving our site:
func customFuncPath(r *http.Request) string {
	path := r.URL.Path
	if path == "/mysite" || path == "/mysite/" {
		path = "/mysite/intro"
	}
	return path[len("/mysite/"):]
}
