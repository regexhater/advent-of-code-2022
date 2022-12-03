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
	var totalScoreForPart1 int
	var totalScoreForPart2 int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			letters := strings.Split(line, " ")
			if len(letters) != 2 {
				log.Fatalf("malformed input, expected two elements per line but got %v. Line: %v\n", len(letters), line)
			}
			oponentPlay := newPlayWithCode(letters[0])
			yourPlayForPart1 := newPlayWithCode(letters[1])
			yourPlayForPart2 := newPlayWithSign(getRPSForResult(oponentPlay, letters[1]))
			totalScoreForPart1 += roundResult(*oponentPlay, *yourPlayForPart1)
			totalScoreForPart2 += roundResult(*oponentPlay, *yourPlayForPart2)
		}
	}
	log.Printf("---Total score---\nPart 1: %v points\nPart 2: %v Points\n", totalScoreForPart1, totalScoreForPart2)
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

func getRPSForResult(oponentPlay *Play, result string) rps {
	var rpsForWantedResult rps
	switch result {
	//win
	case "Z":
		rpsForWantedResult = oponentPlay.Weakness
	// draw
	case "Y":
		rpsForWantedResult = oponentPlay.Played
	// lose
	case "X":
		{
			if oponentPlay.Played == Rock {
				rpsForWantedResult = Scissors
			} else if oponentPlay.Played == Paper {
				rpsForWantedResult = Rock
			} else {
				rpsForWantedResult = Paper
			}
		}
	}

	return rpsForWantedResult
}
