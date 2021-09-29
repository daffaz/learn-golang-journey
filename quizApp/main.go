package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func exitWithMessage(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func parseLines(someSlice [][]string) []problem {
	var returnedSlices []problem = make([]problem, len(someSlice))
	for i, value := range someSlice {
		returnedSlices[i] = problem{
			question: value[0],
			answer:   value[1],
		}
	}

	return returnedSlices
}

// [
// 	{question: 1, answer: 2},
// 	{question: 1, answer: 2},
// 	{question: 1, answer: 2},
// 	{question: 1, answer: 2},
// ]

func main() {
	var csvFileName *string = flag.String("csv", "problems.csv", "a csv file in the format question,answer")
	flag.Parse()

	var file, err = os.Open(*csvFileName)

	// If failed to open CSV file
	if err != nil {
		exitWithMessage(fmt.Sprintf("Failed to open CSV file: %s", *csvFileName))
	}

	var readFile *csv.Reader = csv.NewReader(file)
	parsedCsvFile, err := readFile.ReadAll()

}
