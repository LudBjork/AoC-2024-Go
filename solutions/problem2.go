package solutions

import (
	"fmt"
	"ludbjork/aoc-2024/commons"
	"strings"
)

func SolveProblem2() {
	input := commons.ReadInput("inputs/problem_2.txt")

	p2_part2(input)
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

func p2_part2(input []string) {

	reports := make(map[int][]int)

	SEPARATOR := " "

	for index, value := range input {
		reports[index] = commons.StringSliceToIntSlice(strings.Split(value, SEPARATOR))
	}

	count := 0
	for _, report := range reports {
		if isReportSafeV2(report, 0) {
			count++
		}
	}

	fmt.Println(count)
}

func isReportSafe(report []int) bool {
	var strictInc bool
	var strictDec bool
	for i, _ := range report {
		if i > 0 {
			if i == 1 {
				strictInc = report[i] > report[i-1]
				strictDec = report[i] < report[i-1]
			}

			strictInc = strictInc && (report[i] > report[i-1])
			strictDec = strictDec && (report[i] < report[i-1])

			if commons.ComputeDistance(report[i], report[i-1]) == 0 {
				return false
			}
			if commons.ComputeDistance(report[i], report[i-1]) > 3 {
				return false
			}
		}
	}
	if !strictInc && !strictDec {
		return false
	}

	return true
}

func isReportSafeV2(report []int, retryAttempts int) bool {
	// Credit where credit is due: https://github.com/proxyvix/AoC_2024/blob/master/day2/day2.go
	// I didn't get it until reading their answer
	if isReportSafe(report) {
		return true
	}
	for i, _ := range report {
		var temp []int
		temp = append(temp, report[:i]...)
		temp = append(temp, report[i+1:]...)
		if isReportSafe(temp) {
			return true
		}
	}
	return false
}
