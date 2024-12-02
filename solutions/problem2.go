package solutions

import (
	"fmt"
	"ludbjork/aoc-2024/commons"
	"strings"
)

func SolveProblem2() {
	input := commons.ReadInput("inputs/problem_2.txt")

	p2_part1(input)
}

func p2_part1(input []string) {

	reports := make(map[int][]int)

	SEPARATOR := " "

	for index, value := range input {
		reports[index] = commons.StringSliceToIntSlice(strings.Split(value, SEPARATOR))
	}

	count := 0
	for _, report := range reports {
		if isReportSafe(report) {
			count++
		}
	}

	fmt.Println(count)
}

func isReportSafe(report []int) bool {
	for i, _ := range report {
		if i > 0 {
			if commons.ComputeDistance(report[i], report[i-1]) == 0 {
				return false
			}
			if commons.ComputeDistance(report[i], report[i-1]) > 3 {
				return false
			}
			if (report[i] > report[0]) != (report[i-1] > report[0]) {
				return false
			}
		}
	}
	return true
}
