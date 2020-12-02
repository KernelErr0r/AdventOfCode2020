package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	defer file.Close()

	if err != nil {
		panic(err)
	}

	count := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		lowerBound, err := strconv.Atoi(before(line, "-"))

		if err != nil {
			panic(err)
		}

		higherBound, err := strconv.Atoi(before(after(line, "-"), " "))

		if err != nil {
			panic(err)
		}

		character := after(before(line, ":"), " ")
		value := after(line, ": ")

		characterCount := strings.Count(value, character)

		if characterCount >= lowerBound && characterCount <= higherBound {
			count++
		}
	}

	fmt.Println(count)
}

func before(value string, valueToFind string) string {
	return strings.Split(value, valueToFind)[0]
}

func after(value string, valueToFind string) string {
	return strings.Split(value, valueToFind)[1]
}