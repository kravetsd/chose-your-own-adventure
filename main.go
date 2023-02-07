package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Intro struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type NewYork struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Debate struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type SeanKelly struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type MarkBates struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Denver struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Home struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []interface{} `json:"options"`
}

type Book struct {
	Intro     Intro     `json:"intro"`
	NewYork   NewYork   `json:"new-york"`
	Debate    Debate    `json:"debate"`
	SeanKelly SeanKelly `json:"sean-kelly"`
	MarkBates MarkBates `json:"mark-bates"`
	Denver    Denver    `json:"denver"`
	Home      Home      `json:"home"`
}

func main() {
	fmt.Println("Hello, cyoa!")
	fl, err := os.Open("gopher.json")
	if err != nil {
		log.Fatal("Openning file:", err)
	}
	defer fl.Close()

	bk := new(Book)
	jsondec := json.NewDecoder(fl)
	err = jsondec.Decode(bk)
	if err != nil {
		log.Fatal("Error decoding json:", err)
	}

	mux := defaultMux()
	http.ListenAndServe(":8080", bookHandler(bk, mux))

}

func bookHandler(bk *Book, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/intro" {
			fmt.Fprintln(w, bk.Intro)
		}
		log.Default().Println("Redirecting to fallback")
		fallback.ServeHTTP(w, r)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
