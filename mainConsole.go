package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)


func main() {
	lines, err := readLines("./words.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}

	stringList := make(map[string]int)
	var keys []string
	for _, line := range lines {
		lowercaseWord := strings.ToLower(line)
		processedString := reg.ReplaceAllString(lowercaseWord,"")
		splitString := strings.Split(processedString, " ")
		for _, stringWord := range splitString {
			_, ok := stringList[stringWord]
			if !(ok) {
					if(len(stringWord)>0){
						keys = append(keys, stringWord)
						stringList[stringWord] = 1
					}

			} else {
				stringList[stringWord] += 1
			}
		}
		//s := []string{line, "is", "a", "joined", "string\n"};
		//fmt.Println(i, strings.Join(s, " "))
	}

	sort.Strings(keys)

	for _, value := range keys {
		fmt.Println("W: ", value, ":", stringList[value])
	}


}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[word]++
	}
	return counts
}
