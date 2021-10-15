package cerita

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
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
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tmplt.Execute(w, chapter)
		if err != nil {
			log.Printf("%+v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
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
<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
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
