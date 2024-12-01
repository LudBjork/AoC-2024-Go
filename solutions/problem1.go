package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput("./inputs/problem_1.txt")
	// _, str because first element in range is the index
	var left []int
	var right []int
	// See problem_1.txt for why SEPARATOR is like this
	SEPARATOR := "   "
	for _, str := range input {
		left = append(left, SimpleStringToInt(strings.Split(str, SEPARATOR)[0]))
		right = append(right, SimpleStringToInt(strings.Split(str, SEPARATOR)[1]))
	}

	// sort both arrays in-place
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})
	var distances []int
	for index, _ := range left {
		distances = append(distances, ComputeDistance(left[index], right[index]))
	}
	sum := 0
	for _, partial := range distances {
		sum += partial
	}
	fmt.Println(sum)
}

func ComputeDistance(left int, right int) int {
	if left > right {
		return left - right
	}
	return right - left
}
func SimpleStringToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

// Returns every line of input file as a slice.
// ENSURE INPUT HAS NO BLANK LINE
func ReadInput(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	file.Close()
	return input
}
