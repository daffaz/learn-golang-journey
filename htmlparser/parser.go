package htmlparser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func linkNode(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var res []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, linkNode(c)...)
	}
	return res
}

func buildText(link *html.Node) string {
	if link.Type == html.TextNode {
		return link.Data
	}
	if link.Type != html.ElementNode {
		return ""
	}
	var res string
	for c := link.FirstChild; c != nil; c = c.NextSibling {
		res += buildText(c) + " "
	}
	return strings.Join(strings.Fields(res), " ")
}

func buildLink(link *html.Node) Link {
	var res Link
	for _, node := range link.Attr {
		if node.Key == "href" {
			res.Href = node.Val
			break
		}
	}
	res.Text = buildText(link)
	return res
}

func Parser(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Printf("error, %v", err)
	}
	nodes := linkNode(doc)
	var someLink []Link
	for _, node := range nodes {
		someLink = append(someLink, buildLink(node))
	}
	return someLink, nil
}
