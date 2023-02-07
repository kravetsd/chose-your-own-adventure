package cyoa

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
