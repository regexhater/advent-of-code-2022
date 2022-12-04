package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func main() {
	log.Println("Advent of Code 2022 day three")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file\n%v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sumOfPriorities int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			duplicate, err := findItemInBothComp(line)
			if err != nil {
				log.Fatalf("Error while finding duplicate: %v\n", err)
			}
			sumOfPriorities += calculatePriority(duplicate)
		}
	}

	log.Printf("---Result---\nSum of priorioties: %v\n", sumOfPriorities)
}

func findItemInBothComp(line string) (rune, error) {
	rucksackSize := len(line) // 4
	compSplit := rucksackSize / 2
	items := []rune(line)
	visitedItemsInLeftComp := make(map[rune]int)
	visitedItemsInRightComp := make(map[rune]int)
	for i := 0; i < compSplit; i++ {
		j := compSplit + i
		leftItem := items[i]
		rightItem := items[j]
		if leftItem == rightItem {
			return leftItem, nil
		}
		if elementExist(visitedItemsInRightComp, leftItem) {
			return leftItem, nil
		}
		visitedItemsInLeftComp[leftItem] = 1
		if elementExist(visitedItemsInLeftComp, rightItem) {
			return rightItem, nil
		}
		visitedItemsInRightComp[rightItem] = 1
	}
	return 0, errors.New("could not find an item in both comp")
}

func elementExist(items map[rune]int, element rune) bool {
	_, exists := items[element]
	return exists
}

func calculatePriority(letter rune) int {
	asci := int(letter)
	if asci < 64 || asci > 122 {
		log.Fatalf("Bad rucksack input detected: %v\n,", letter)
	}
	var normalizedPriority int
	if asci > 96 {
		normalizedPriority = asci - 96
	} else {
		normalizedPriority = asci - 37
	}
	return normalizedPriority
}
