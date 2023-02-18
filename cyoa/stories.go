package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

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
	tp    *template.Template
}

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

	//The code bellow shows that calls can be chained together.
	// var err error
	// sh.tp, err = template.New("story.html").Funcs(myFuncMap).ParseFiles("templates/story.html")
	// if err != nil {
	// 	log.Fatal("Parsing template:", err)
	// }

	// we trim leading "/" in path
	path := r.URL.Path

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := sh.Story[path]; ok {
		err := sh.tp.Execute(w, chapter)
		if err != nil {
			log.Fatal("Executing template:", err)
			http.Error(w, "Something wenr worng", http.StatusInternalServerError)
		}
	} else {
		log.Default().Printf(" %s does not exist", path)
		http.Error(w, "This chapter does not exist.", http.StatusNotFound)
	}
}

//Why we return a pointer to a StoryHandler? what is a practice in this case?
func NewStoryHandler(st Story) *StoryHandler {
	// The code bellow add some debugging just for study purposes.
	var err error
	// template.New function is used because we want to add a custom function to a template before it is parsed.
	// template name must be the a base name of the file that is passed to a template.ParseFiles function below.!!!
	tp := template.New("story.html")
	// fmt.Printf("sh.tp: %+#v", sh.tp)

	// template.Funcs function is used to add a custom function to a template.
	tp = tp.Funcs(myFuncMap)
	// fmt.Printf("ADDED FUNCS TO sh.tp: %+v", sh.tp)

	// template.ParseFiles function is used to parse a template from a file.
	tp, err = tp.ParseFiles("templates/story.html")

	if err != nil {
		log.Printf("Parsing template: %v", err)
	}
	return &StoryHandler{Story: st, tp: tp}
}
