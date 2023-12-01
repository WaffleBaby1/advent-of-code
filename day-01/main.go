package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// function read the file into memory line by line
// all file contents is stores in lines[]
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

func splitAfter(word string, re *regexp.Regexp) (r []string) {
	re.ReplaceAllStringFunc(word, func(x string) string {
		word = strings.Replace(word, x, "::"+x, -1)
		return word
	})
	for _, x := range strings.Split(word, "::") {
		if x != "" {
			r = append(r, x)
		}
	}
	return
}

func getValueOfString(word string) (numValue int) {
	value := string(word[0])
	fmt.Println(value)
	numValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	fmt.Println("ending function")
	return numValue
}

func main() {
	//open the file for reading
	var totalValue int
	lines, err := readLines("stars.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i := range lines {
		wordToParse := lines[i]
		parsedWord := splitAfter(wordToParse, regexp.MustCompile("[0-9]"))
		fmt.Println(parsedWord)

		splitLength := len(parsedWord)

		if splitLength == 1 {
			re := regexp.MustCompile("[0-9]")
			soloWord := parsedWord[0]
			hasNumber := re.Match([]byte(soloWord))
			if hasNumber {
				intValue := getValueOfString(soloWord)
				intValue *= 11
				fmt.Printf("the value of word is: %d\n", intValue)
				totalValue += intValue
			}
		} else {
			for j := range parsedWord {
				word := parsedWord[j]
				firstLetter := string(word[0])
				re := regexp.MustCompile("[0-9]")
				hasNumber := re.Match([]byte(firstLetter))
				if hasNumber {
					lastNumber := string(parsedWord[splitLength-1][0])
					wholeNumber := firstLetter + lastNumber
					fmt.Println(wholeNumber)
					numValue, err := strconv.Atoi(wholeNumber)
					if err != nil {
						panic(err)
					}
					totalValue += numValue
					break
				}
			}
		}
		fmt.Printf("The total value is : %d\n", totalValue)
	}
	fmt.Printf("The total value is : %d\n", totalValue)
}
