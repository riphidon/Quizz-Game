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
	q, err := os.Open("quizz.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(q))

	for {
		record, err := r.Read()
		if err == io.EOF {
			if count > 10 {
				fmt.Println("You Won !!!")
			} else {
				fmt.Println("Try again looser")
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
