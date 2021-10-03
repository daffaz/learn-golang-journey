package main

import (
	"cerita"
	"encoding/json"
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

	fileInJson := json.NewDecoder(fileStream)
	var cerita cerita.Story

	if err = fileInJson.Decode(&cerita); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", cerita)
}
