package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arl/statsviz"
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

	// mux := http.NewServeMux()
	// statsviz.Register(mux)

	// go func() {
	// 	fmt.Println("Visit me at http://localhost:6060/debug/statsviz/")
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	log.Println("Visit performance tool at http://localhost:6060/debug/statsviz/")
	statsviz.RegisterDefault()
	http.ListenAndServe(":8080", nil)

}
