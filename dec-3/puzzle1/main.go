package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cast"
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

	columnSums, rows := sumEachColumn(scanner)
	fmt.Printf("The column sums are: %v. %v rows\n", columnSums, rows)
	gamma, epsilon := determineBinaryValues(columnSums, rows)
	fmt.Printf("The hex values are: %v and %v. \n", gamma, epsilon)
	final := convertValsToDecimalandCompute(gamma, epsilon)

	fmt.Printf("The answer is %v\n.", final)
}

func sumEachColumn(scanner *bufio.Scanner) ([12]int, float32) {
	var results [12]int
	var rows float32 = 0.0

	for scanner.Scan() {
		rows += 1.0
		line := scanner.Text()
		fmt.Println(line)
		for i := 0; i < len(line); i++ {
			// fmt.Printf("value %d.\n", cast.ToInt(line[i]))
			results[i] += cast.ToInt(line[i]) - 48
			// fmt.Println(line[0])
		}
	}

	return results, rows
}

func determineBinaryValues(values [12]int, rows float32) (string, string) {
	var gamma, epsilon bytes.Buffer
	for i := 0; i < 12; i++ {
		dec := float32(values[i]) / rows
		if dec > .5 {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}

	return gamma.String(), epsilon.String()
}

func convertValsToDecimalandCompute(gamma string, epsilon string) int {
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		fmt.Printf("Error converting gamma binary to decimal: %v.\n", err)
	}

	epsInt, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		fmt.Printf("Error converting epsilon binary to decimal: %v.\n", err)
	}

	return int(gammaInt) * int(epsInt)
}
