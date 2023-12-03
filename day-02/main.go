package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"regexp"
	"strconv"
	"unicode"
)

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

func checkColon(word string) (bool){
	for i := 0; i < len(word); i++ {
		if string(word[i]) == ";"{
			return true
		}
	}
	return false
}

func gameValues(gameSplit []string)(int, int, int){
	red := regexp.MustCompile(`red.?`)
	blue := regexp.MustCompile(`blue.?`)
	green := regexp.MustCompile(`green.?`)
	var redPulled int
	var bluePulled int
	var greenPulled int
	for i := range gameSplit {
			var redNum int
			var blueNum int
			var greenNum int
			var stringValue string
			spacePull := gameSplit[i]
		    pull := strings.ReplaceAll(spacePull, " ", "")
			for _, char := range pull {
				if unicode.IsDigit(char) {
					stringValue += string(char)
				}
			}
			if red.Match([]byte(pull)) {
				numValue, err := strconv.Atoi(string(stringValue))
				if err != nil {
						panic(err)
				}
				redNum = numValue
			} else if  blue.Match([]byte(pull)){
				numValue, err := strconv.Atoi(string(stringValue))
				if err != nil {
						panic(err)
				}
				blueNum = numValue
			}else if green.Match([]byte(pull)){
				numValue, err := strconv.Atoi(string(stringValue))
				if err != nil {
						panic(err)
				}
				greenNum = numValue
			}
			redPulled += redNum
			bluePulled += blueNum
			greenPulled += greenNum
		}
	return redPulled, bluePulled, greenPulled
}

func main() {
	
	redPossible := 12
	greenPossible := 13
	bluePossible := 14
	var totalPossible int

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for line := range lines {
		redPulled := 0
		greenPulled := 0
		bluePulled := 0
		possible := false
		myFlag := false

		gameNum := strings.SplitAfter(lines[line], ":")[0]
		fmt.Printf("Game Num is %s\n",gameNum)
		game := strings.SplitAfter(lines[line], ":")[1]
		// fmt.Println(game)
		commaSplit := strings.Split(game, ",")
		var colonGame []string
		for j := range commaSplit {
			if checkColon(commaSplit[j]) {
				semiSplit := strings.Split(commaSplit[j], ";")
				for k := range semiSplit {
					colonGame = append(colonGame, semiSplit[k])
				}
			} else {
				colonGame = append(colonGame, commaSplit[j])
			}
		}

		if myFlag {
			fmt.Println(colonGame)
			redPulled, bluePulled, greenPulled = gameValues(colonGame)
			fmt.Println(redPulled, " ", bluePulled, " ", greenPulled)
			if redPulled <= redPossible && bluePulled <= bluePossible && greenPulled <= greenPossible {
			 	fmt.Printf("Game %d is possible\n", line+1)
				totalPossible += (line+1)
			}
		} else {
			colonSplit := strings.Split(game, ";")
			for i := range colonSplit {
				// fmt.Println(colonSplit[i])
				appendGame := strings.Split(colonSplit[i], ",")
				redPulled, bluePulled, greenPulled = gameValues(appendGame)
				fmt.Println(redPulled, " ", bluePulled, " ", greenPulled)
				if redPulled <= redPossible && bluePulled <= bluePossible && greenPulled <= greenPossible {
			 		fmt.Printf("Game %d is possible\n", line+1)
					possible = true
				} else {
					possible = false
					fmt.Println("Game is Impossible")
					break
				}
			}	
		}
	if possible{
		totalPossible += (line+1)
	}
	}
	 fmt.Printf("Total possible is: %d\n", totalPossible)
}