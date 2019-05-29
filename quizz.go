package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var res string
	var count int
	var total = getNbQuestions()
	var file = readFile()
	fmt.Println(total)
	for {
		record, err := file.Read()
		if err == io.EOF {
			if count > 10 {
				fmt.Printf("You scored %v out of %v, good job\n", count, total)
			} else {
				fmt.Printf("You scored %v out of %v, try again\n", count, total)
			}
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("What %v is equal to?\n", record[0])
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			res = scanner.Text()
		}
		if res != record[1] {
			fmt.Printf("Wrong Answer!! %v equals %s\n", record[0], record[1])
		} else {
			count++
		}

	}

}

func readFile() csv.Reader {
	q, err := os.Open("quizz.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(q))
	return *r
}

func getNbQuestions() int {
	var file = readFile()
	questions, err := file.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return len(questions)
}
