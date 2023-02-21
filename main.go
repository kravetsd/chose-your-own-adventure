package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
		for {
			if <-ch == "start" {
				currentStory := story["intro"]
				fmt.Println(currentStory.Title)
				fmt.Println("------------")
				fmt.Println(strings.Join(currentStory.Story, " "))
				fmt.Println("------------")
				fmt.Println("Press the number to choose how to proceed with this story : ")
				for i, st := range currentStory.Options {
					fmt.Printf("%d.  %v \n", i+1, st.Text)
				}
				fmt.Println("------------")
				var done string
				fmt.Println("Please enter \"done\" to exit the story... ")
				fmt.Scan(&done)
				//ch <- "done"
				ch <- done

			}
		}
	}()

	// orchestaating goroutines:
	for {
		sig := <-ch
		switch sig {
		case "done":
			fmt.Println("Bye-bye!")
			return

		default:
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
