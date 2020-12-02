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

		firstIndex, err := strconv.Atoi(before(line, "-"))

		if err != nil {
			panic(err)
		}

		firstIndex--

		secondIndex, err := strconv.Atoi(before(after(line, "-"), " "))

		if err != nil {
			panic(err)
		}

		secondIndex--

		character := after(before(line, ":"), " ")
		value := after(line, ": ")

		isFirstCharacterValid := string([]rune(value)[firstIndex]) == character
		isSecondCharacterValid := string([]rune(value)[secondIndex]) == character

		if (isFirstCharacterValid || isSecondCharacterValid) && !(isFirstCharacterValid && isSecondCharacterValid) {
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