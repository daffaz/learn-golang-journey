package main

import (
	"cerita"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("file", "cerita.json", "Add a file for your cerita web")
	flag.Parse()

	// var
	fileStream, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	fileInJson, err := cerita.StreamToJson(fileStream)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", fileInJson)
}
