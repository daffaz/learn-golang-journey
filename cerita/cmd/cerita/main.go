package main

import (
	"cerita"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	const PORT = 3000

	fileName := flag.String("file", "cerita.json", "Add a file for your cerita web")
	flag.Parse()

	fileStream, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	fileInJson, err := cerita.StreamToJson(fileStream)

	if err != nil {
		panic(err)
	}
	h := cerita.NewHandler(fileInJson)
	fmt.Println("Starting in port :", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), h))
}
