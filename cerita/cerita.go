package cerita

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

var tmplt *template.Template

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

type handlr struct {
	s Story
}

func init() {
	tmplt = template.Must(template.New("").Parse(HTMLTemplate))
}

func StreamToJson(reader io.Reader) (Story, error) {
	var fileInJson = json.NewDecoder(reader)
	var story Story
	if err := fileInJson.Decode(&story); err != nil {
		panic(err)
	}

	return story, nil
}

func NewHandler(s Story) http.Handler {
	return handlr{s}
}

func (h handlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tmplt.Execute(w, h.s["intro"])
	if err != nil {
		log.Println(err)
	}
}

var HTMLTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Pilih cerita mu sendiri | {{.Title}}</title>
</head>
<body>
    <section class="page">
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
                <li><a href="/{{.Arc}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </section>
</body>
</html>
`
