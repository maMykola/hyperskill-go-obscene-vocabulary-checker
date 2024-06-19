package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var filename string

	fmt.Scan(&filename)

	tabooWords := loadTaboo(filename)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		sentence := scanner.Text()
		if strings.EqualFold(sentence, "exit") {
			break
		}

		hasPeriod := len(sentence) > 0 && sentence[len(sentence)-1] == '.'
		words := strings.Fields(strings.TrimRight(sentence, "."))
		for i, word := range words {
			if isTaboo(word, &tabooWords) {
				words[i] = strings.Repeat("*", len(word))
			}
		}

		fmt.Print(strings.Join(words, " "))
		if hasPeriod {
			fmt.Print(".")
		}

		fmt.Println()
	}

	fmt.Println("Bye!")
}

func loadTaboo(filename string) map[string]interface{} {
	var table = make(map[string]interface{})

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		table[strings.ToLower(scanner.Text())] = nil
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return table
}

func isTaboo(word string, table *map[string]interface{}) bool {
	_, ok := (*table)[strings.ToLower(word)]
	return ok
}
