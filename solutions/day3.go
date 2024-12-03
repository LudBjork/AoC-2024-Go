package solutions

import (
	"fmt"
	"ludbjork/aoc-2024/commons"
	"regexp"
	"strings"
)

func SolveProblem3() {
	input := commons.ReadInputV2("inputs/day3.txt")
	p3_part1(input)

	inputAsStr := commons.ReadInputV2("inputs/day3.txt")
	p3_part2(inputAsStr)
}

func p3_part1(input string) {
	commandRegEx := regexp.MustCompile("(mul\\([, [0-9]+\\))")
	multiplesSum := 0
	// This code is dreadful but solves the problem :)

	found := commandRegEx.FindAllString(input, -1)
	if len(found) > 0 {
		for _, mulStr := range found {
			temp := strings.Split(mulStr, " ")
			for _, mul := range performMultiplication(temp) {
				multiplesSum += mul
			}
		}
	}
	fmt.Println(multiplesSum)
}

func p3_part2(input string) {
	ignoreRegex := regexp.MustCompile("don't\\(\\)(.*?)do\\(\\)")
	commandRegEx := regexp.MustCompile("(mul\\([, [0-9]+\\))")

	cleanedUpInput := ignoreRegex.ReplaceAllString(input, "")

	productsArr := performMultiplication(commandRegEx.FindAllString(cleanedUpInput, -1))
	sum := 0
	for _, product := range productsArr {
		sum += product
	}
	fmt.Println(sum)
}
func performMultiplication(pairArr []string) []int {
	r := regexp.MustCompile("[0-9]+")
	var products []int
	for _, mulStr := range pairArr {
		temp := commons.StringSliceToIntSlice(r.FindAllString(mulStr, 2))
		products = append(products, temp[0]*temp[1])
	}
	return products
}
