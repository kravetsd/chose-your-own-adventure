package cyoa

import (
	"encoding/json"
	"io"
	"log"
)

type Option struct {
	Text  string `json:"text"`
	Title string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
	Path    string
}

type Story map[string]Chapter

func JsonStory(r io.Reader) (Story, error) {
	st := make(Story)
	sj := json.NewDecoder(r)
	if err := sj.Decode(&st); err != nil {
		log.Fatal("Error decoding json:", err)
		return nil, err
	}
	return st, nil

}

func JsonStoryDecode(bt []byte) (Story, error) {
	st := make(Story)
	err := json.Unmarshal(bt, &st)
	if err != nil {
		log.Println("Error unmarshalling json: ", err)
	}
	return st, err
}
