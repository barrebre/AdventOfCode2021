package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Aim   int
	Depth int
}

// Use an enum to track the directions more easily
type Direction int

const (
	Undefined Direction = iota
	Forward
	Up
	Down
)

func main() {
	// Open the file of inputs, which I named input.txt
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file: %v\n.", filename)
		os.Exit(1)
	}

	// Instantiate a scanner on the opened file to read in line-by-line
	scanner := bufio.NewScanner(f)

	pos := Position{0, 0}
	distanceTraveled := 0

	// Iterate through the rest of the file
	for scanner.Scan() {
		line := scanner.Text()
		direction, distance := parseTravel(line)

		switch direction {
		case Forward:
			distanceTraveled += distance
			pos.Depth += distance * pos.Aim
		case Up:
			pos.Aim -= distance
		case Down:
			pos.Aim += distance
		}
	}

	ans := pos.Depth * distanceTraveled
	fmt.Printf("The final answer is %v.\n", ans)
}

func parseTravel(line string) (Direction, int) {
	split := strings.Split(line, " ")

	distance, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Printf("Encountered error parsing int in string: %v\n.", err)
	}

	switch split[0] {
	case "forward":
		return Forward, distance
	case "down":
		return Down, distance
	default:
		return Up, distance
	}
}
