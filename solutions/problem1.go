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
	left, right := ReadInput("./inputs/problem_1.txt")
	//part1(left, right)
	part2(left, right)
}

func part1(left []int, right []int) {

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

func part2(left []int, right []int) {
	similarityMap := CreateSimilarityMap(left, right)
	sum := 0
	for _, value := range left {
		sum += value * similarityMap[value]
	}
	fmt.Println(sum)
}

func CreateSimilarityMap(left []int, right []int) map[int]int {
	var similarityMap map[int]int
	similarityMap = make(map[int]int)

	for _, valueLeft := range left {
		// check if already iterated to minimise repeats
		if similarityMap[valueLeft] == 0 {
			for _, valueRight := range right {
				if valueLeft == valueRight {
					similarityMap[valueLeft]++
				}
			}
		}
	}
	return similarityMap
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
func ReadInput(filepath string) ([]int, []int) {
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

	var left []int
	var right []int
	// See problem_1.txt for why SEPARATOR is like this
	SEPARATOR := "   "
	for _, str := range input {
		left = append(left, SimpleStringToInt(strings.Split(str, SEPARATOR)[0]))
		right = append(right, SimpleStringToInt(strings.Split(str, SEPARATOR)[1]))
	}

	return left, right
}
