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

	ch := make(chan string, 2)
	go func() {
		sh := cyoa.NewStoryHandler(story, cyoa.WithTemplatePath("templates/story_new.html"), cyoa.WithUrlPath(customFuncPath))
		statsviz.RegisterDefault()
		http.Handle("/mysite/", sh)
		log.Println("\n------------\nVisit performance tool at http://localhost:8080/debug/statsviz/\nYou also can read this story via web interface at http://localhost:8080/mysite/\n------------")
		ch <- "start"
		http.ListenAndServe(":8080", nil)

	}()
	go func() {
		if <-ch == "start" {
			fmt.Println("Starting CLI story")
			fmt.Println("Please enter \"done\" to exit the story... ")
			var done string
			fmt.Scan(&done)
			ch <- "done"
		} else {
			fmt.Println(<-ch)
		}
	}()

	// orchestaating goroutines:
	for {
		sig := <-ch
		switch sig {
		case "done":
			fmt.Println("Bye-bye!")
			return

		case "start":
			ch <- "start"
		}
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
