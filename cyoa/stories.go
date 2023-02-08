package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func JsonStory(r io.Reader) (Story, error) {
	st := make(Story)
	sj := json.NewDecoder(r)
	if err := sj.Decode(&st); err != nil {
		log.Fatal("Error decoding json:", err)
		return nil, err
	}
	return st, nil

}

func (s Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world from story!"))
}

type Option struct {
	Text  string `json:"text"`
	Title string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Chapter
