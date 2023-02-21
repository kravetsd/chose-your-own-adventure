package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
		ch <- "start"
		http.ListenAndServe(":8080", nil)
	}()

	if <-ch == "start" {
		go func() {
			currentStory := story["intro"]
			for {
				cyoa.ShowStoryCli(currentStory)
				var done string
				fmt.Println("Please enter \"done\" to exit the story... ")
				fmt.Scan(&done)
				if done == "done" {
					ch <- done
					break
				}
				choice, err := strconv.Atoi(done)
				if err != nil {
					log.Fatal("Erro choosing option:", err)
				}

				cs, ok := story[currentStory.Options[choice-1].Title]
				if !ok {
					fmt.Println("Sorry. Incorrect input. Please choose one of the options above")
					continue
				}
				currentStory = cs
			}
		}()
	}

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
