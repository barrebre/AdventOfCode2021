package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Start increases at -1, because the logic will automatically add 1 in the first check
	increases := -1

	// Open the file of inputs, which I named input.txt
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file: %v\n.", filename)
		os.Exit(1)
	}

	// Instantiate a scanner on the opened file to read in line-by-line
	scanner := bufio.NewScanner(f)

	// Make an array of 3 empty elements
	values := make([]int, 3)

	// make the var we'll store the read string into
	var line string

	for i := 0; i < 2; i++ {
		// Invoke Scan to get the first two lines into memory
		scanner.Scan()
		line = scanner.Text()

		// Conver the string to an int
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Couldn't convert line into Int: %v.\n", line)
			os.Exit(1)
		}

		// Trim the first value in the array, and then add the new value to the end
		values = values[1:]
		values = append(values, lineInt)
	}

	// This is how we'll keep track of the current sum vs the previous one
	lastSum := 0
	sum := 0

	// Iterate through the rest of the file
	for scanner.Scan() {
		line = scanner.Text()

		lineInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Couldn't convert line into Int: %v.\n", line)
			os.Exit(1)
		}

		// Trim the first value in the array, and then add the new value to the end
		values = values[1:]
		values = append(values, lineInt)

		// pass the object references to the function so we can compute without more alloc's
		sum = checkValues(&values)

		if sum > lastSum {
			increases++
		}

		// replace the old lastSum with the current one
		lastSum = sum
	}

	fmt.Printf("There were %v increments.\n", increases)
}

// checkValues returns the sum of values in an int array
func checkValues(values *[]int) int {
	sum := 0
	for _, value := range *values {
		sum += value
	}

	return sum
}
