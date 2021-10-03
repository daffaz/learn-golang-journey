package cerita

import (
	"encoding/json"
	"io"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Story map[string]Chapter

func StreamToJson(reader io.Reader) (Story, error) {
	var fileInJson = json.NewDecoder(reader)
	var story Story
	if err := fileInJson.Decode(&story); err != nil {
		panic(err)
	}

	return story, nil
}
