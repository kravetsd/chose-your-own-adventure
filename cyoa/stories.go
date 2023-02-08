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

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

type StoryHandler struct {
	Story Story
}

//Why we return a pointer to a StoryHandler? what is a practice in this case?
func NewStoryHandler(st Story) *StoryHandler {
	return &StoryHandler{Story: st}
}
