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
		records, err := r.Read()
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
		fmt.Printf("What %v is equal to?\n", records[0])
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			res = scanner.Text()
		}
		if res != records[1] {
			fmt.Printf("Wrong Answer!! %v equals %s\n", records[0], records[1])
		} else {
			count++
		}

	}

}
