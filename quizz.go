package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var res string
	var count int
	var wg = sync.WaitGroup{}
	var total = getNbQuestions()
	var file = readFile()

	ch := make(chan string, total)
	wg.Add(2)

	fmt.Println("Press start")
	gameStart := bufio.NewScanner(os.Stdin)
	if gameStart.Scan() {
		timer := time.AfterFunc(10*time.Second, func() {
			fmt.Println("\nTime is over")
			ch <- fmt.Sprintf("You scored %v out of %v\n", count, total)
			close(ch)
		})
		go func(ch <-chan string) {

			for {
				if i, ok := <-ch; ok {
					fmt.Println(i)
				} else {
					break
				}
			}
			wg.Done()
		}(ch)

		go func(ch chan<- string) {

			defer timer.Stop()
			for {
				record, err := file.Read()
				if err == io.EOF {
					ch <- fmt.Sprintf("You scored %v out of %v\n", count, total)
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("What %v is equal to?\n", record[0])
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					res = scanner.Text()
					if res != record[1] {
						ch <- fmt.Sprintf("Wrong Answer!! %v equals %s\n", record[0], record[1])
					} else {
						ch <- fmt.Sprintf("Correct %v equals %s\n", record[0], record[1])
						count++
					}
				}
			}

			wg.Done()
		}(ch)
		wg.Wait()
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
