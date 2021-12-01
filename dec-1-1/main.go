package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	increases := 0

	// Open the file of inputs, which I named input.txt
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file: %v\n.", filename)
		os.Exit(1)
	}

	// Instantiate a scanner on the opened file to read in line-by-line
	scanner := bufio.NewScanner(f)

	// Invoke Scan to get the first line into memory
	scanner.Scan()
	prevLine := scanner.Text()

	// Conver the string to an int
	prevLineInt, err := strconv.Atoi(prevLine)
	if err != nil {
		fmt.Printf("Couldn't convert line into Int: %v.\n", prevLine)
		os.Exit(1)
	}

	// Set up the future vars in mem
	var newLine string
	var newLineInt int

	// Iterate through the rest of the file
	for scanner.Scan() {
		newLine = scanner.Text()

		newLineInt, err = strconv.Atoi(newLine)
		if err != nil {
			fmt.Printf("Couldn't convert line into Int: %v.\n", prevLine)
			os.Exit(1)
		}

		// If higher, increment the increases value
		if newLineInt > prevLineInt {
			increases += 1
		}

		// Set the new to old and continue on
		prevLineInt = newLineInt
	}

	fmt.Printf("There were %v increments.\n", increases)
}
