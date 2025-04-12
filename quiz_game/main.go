package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {
	fmt.Println("Loading CSV file... 🐼")
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")

	questions := make([]Quiz, len(data))

	for i, data := range data {
		question := Quiz{
			Question: data[0],
			Answer:   data[1],
		}
		questions[i] = question
	}

	for i, questions := range questions {
		fmt.Printf("Question %v: %v\n", i, questions.Question)
	}
	//fmt.Println(questions)

}

// func loadCsv()
