package main

import (
	"fmt"
	"htmlparser"
	"io/ioutil"
	"strings"
)

func main() {
	exampleHTML, err := ioutil.ReadFile("examples/ex1/example.html")
	if err != nil {
		panic(err)
	}
	r := strings.NewReader(string(exampleHTML))
	links, err := htmlparser.Parser(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)
}
