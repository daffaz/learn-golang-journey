package htmlparser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func dfs(n *html.Node, padding string) {
	// if n.Type == html.ElementNode && n.Data == "a" {
	// 	for _, a := range n.Attr {
	// 		if a.Key == "href" {
	// 		}
	// 	}
	// }
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}

func Parser(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Printf("%v", err)
	}
	dfs(doc, "")
	return nil, nil
}
