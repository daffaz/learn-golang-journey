package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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
			answer:   strings.TrimSpace(value[1]),
		}
	}
	return returnedSlices
}

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

	if err != nil {
		exitWithMessage("Failed to parse the provided CSV")
	}

	var quiz []problem = parseLines(parsedCsvFile)
	var countScore uint16 = 0
	for index, value := range quiz {
		fmt.Printf("Question #%d: %s = ", index+1, value.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == value.answer {
			countScore++
		}
	}
	fmt.Printf("\nYou score is %d out of %d", countScore, len(quiz))
}
