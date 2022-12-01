package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	elfCounter int = 0
)

func main() {
	log.Println("Advent of code 2022 day one")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Cannot read input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	top3Elves := make([]*Elf, 3)
	currentElf := createElf()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			addToTop3IfGreater(top3Elves, currentElf)
			currentElf = createElf()
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not convert %v to int\n", line)
		}
		currentElf.addCalories(calories)
	}
	top3Calories := 0
	rankingMsg := ""
	for i, elf := range(top3Elves) {
		top3Calories += elf.totalCalories
		rankingMsg += fmt.Sprintf("Top %v elf with id: %v and %v calories\n", i + 1, elf.id, elf.totalCalories)
	}
	rankingMsg += fmt.Sprintf("Top 3 elves calories: %v\n", top3Calories)
	log.Printf("--- Ranking ---\n%v", rankingMsg)
}

type Elf struct {
	id int
	totalCalories int
}

func createElf() *Elf {
	elf := Elf {
		elfCounter,
		0,
	}
	elfCounter++
	return &elf
}

func (e *Elf) addCalories(calories int) {
	e.totalCalories += calories
}

func addToTop3IfGreater(top3Elves []*Elf, currentElf *Elf) {
	for i, elf := range(top3Elves) {
		if elf == nil || elf.totalCalories < currentElf.totalCalories {
			moveRankingRight(top3Elves, i)
			top3Elves[i] = currentElf
			return
		}
	}
}

func moveRankingRight(top3Elves []*Elf, index int) {
	for i := len(top3Elves) - 1; i > index; i-- {
		top3Elves[i] = top3Elves[i - 1]
	}
}
