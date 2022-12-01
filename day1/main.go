package main

import (
	"bufio"
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

	topElf := createElf()
	var elves []*Elf
	elves = append(elves, topElf)
	currentElfIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			elves = append(elves, createElf())
			currentElfIndex++
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not convert %v to int\n", line)
		}
		currentElf := elves[currentElfIndex]
		currentElf.addCalories(calories)
		if topElf.totalCalories < currentElf.totalCalories {
			topElf = currentElf
		}
	}
	log.Printf("Elf with most calories %v with number %v\n", topElf.totalCalories, topElf.id + 1)
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
