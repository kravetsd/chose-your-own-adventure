package cyoa

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type StoryHandler struct {
	Story Story
	tp    *template.Template
}

type Opt func(sh *StoryHandler)

func WithPath(path string) Opt {
	fl, err := os.Stat(path)
	if err != nil {
		log.Printf("Error openning custom template file: %v", err)
	}
	return func(sh *StoryHandler) {
		sh.tp, err = template.New(fl.Name()).Funcs(myFuncMap).ParseFiles(path)
		if err != nil {
			log.Printf("Error parsing template %v : %v", path, err)
		}
	}
}

// TODO: Addtional functions can be also handled via functional options later.
var myFuncMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"join": strings.Join,
	"plusOne": func(n int) int {
		n += 1
		return n
	},
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
func NewStoryHandler(st Story, opts ...Opt) *StoryHandler {
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

	sh := &StoryHandler{Story: st, tp: tp}
	for _, opt := range opts {
		opt(sh)
	}
	return sh
}
