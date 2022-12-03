package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("Advent of Code 2022 day two")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open input, %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var totalScore int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			letters := strings.Split(line, " ")
			if len(letters) != 2 {
				log.Fatalf("malformed input, expected two elements per line but got %v. Line: %v\n", len(letters), line)
			}
			oponentPlay := newPlay(letters[0])
			yourPlay := newPlay(letters[1])
			totalScore += roundResult(*oponentPlay, *yourPlay)
		}
	}
	log.Printf("Total score: %v points\n", totalScore)
}

func roundResult(oponentPlay Play, yourPlay Play) int {
	var points int
	switch yourPlay.Played {
	case Rock:
		points += 1
	case Paper:
		points += 2
	case Scissors:
		points += 3
	}

	switch oponentPlay.Played {
	case yourPlay.Weakness:
		points += 0
	case yourPlay.Played:
		points += 3
	default:
		points += 6
	}

	return points
}

