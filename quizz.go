package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var s bool
	var stop bool
	var res string
	var count int
	var total = getNbQuestions()
	var file = readFile()

	fmt.Println("Press start")
	gameStart := bufio.NewScanner(os.Stdin)
	if gameStart.Scan() {
		s = true
	}
	if s != false {
		timer := time.AfterFunc(10*time.Second, func() {
			stop = true
			fmt.Println("\nTime is over")
		})
		defer timer.Stop()
		for {

			record, err := file.Read()
			if stop == true || err == io.EOF {
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

			if stop == false {
				fmt.Printf("What %v is equal to?\n", record[0])
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					res = scanner.Text()
					if res != record[1] {
						fmt.Printf("Wrong Answer!! %v equals %s\n", record[0], record[1])
					} else {
						fmt.Printf("Correct %v equals %s\n", record[0], record[1])
						count++
					}
				}
			}

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
