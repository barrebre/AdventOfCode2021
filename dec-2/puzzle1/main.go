package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Forward int
	Up      int
	Down    int
}

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

	pos := Position{0, 0, 0}

	// Iterate through the rest of the file
	for scanner.Scan() {
		line := scanner.Text()

		direction, distance := parseTravel(line)
		fmt.Println(direction, distance)

		switch direction {
		case Forward:
			pos.Forward += distance
		case Up:
			pos.Up += distance
		case Down:
			pos.Down += distance
		}
	}

	dis := calculateDistance(pos)

	fmt.Printf("The total distance is %v.\n", dis)
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

func calculateDistance(pos Position) int {
	depth := pos.Down - pos.Up

	return depth * pos.Forward
}
