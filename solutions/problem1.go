package solutions

import (
	"fmt"
	"ludbjork/aoc-2024/commons"
	"sort"
	"strings"
)

func SolveProblem1() {
	// DANGER: Depends on where program is ran from. I.e. use project root dir
	// always!!!
	input := commons.ReadInput("inputs/problem_1.txt")
	var left []int
	var right []int
	// See problem_1.txt for why SEPARATOR is like this
	SEPARATOR := "   "
	for _, str := range input {
		left = append(left, commons.SimpleStringToInt(strings.Split(str, SEPARATOR)[0]))
		right = append(right, commons.SimpleStringToInt(strings.Split(str, SEPARATOR)[1]))
	}

	part1(left, right)
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
		distances = append(distances, computeDistance(left[index], right[index]))
	}
	sum := 0
	for _, partial := range distances {
		sum += partial
	}
	fmt.Println(sum)
}

func part2(left []int, right []int) {
	similarityMap := createSimilarityMap(left, right)
	sum := 0
	for _, value := range left {
		sum += value * similarityMap[value]
	}
	fmt.Println(sum)
}

func createSimilarityMap(left []int, right []int) map[int]int {
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

func computeDistance(left int, right int) int {
	if left > right {
		return left - right
	}
	return right - left
}
