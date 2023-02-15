package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
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

var myFuncMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"join": strings.Join,
	"plusOne": func(n int) int {
		n += 1
		return n
	},
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//adding .Funcs(myFuncMap) to a template gives us a custom "join" function inside a template.
	tp, err := template.New("story.html").Funcs(myFuncMap).ParseFiles("templates/story.html")
	if err != nil {
		log.Fatal("Parsing template:", err)
	}

	if chapter, ok := sh.Story[strings.Split(r.URL.Path, "/")[1]]; ok {
		err = tp.Execute(w, chapter)
		if err != nil {
			log.Fatal("Executing template:", err)
		}
	} else {
		err = tp.Execute(w, sh.Story["intro"])
		if err != nil {
			log.Fatal("Executing template:", err)
		}
	}

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
