package solutions

import (
	"fmt"
	"ludbjork/aoc-2024/commons"
	"regexp"
	"strings"
)

func SolveProblem3() {
	input := commons.ReadInput("inputs/day3.txt")
	p3_part1(input)
}

func p3_part1(input []string) {
	regexPattern, _ := regexp.Compile("(mul\\([, [0-9]+\\))")
	multiplesSum := 0
	// This code is dreadful but solves the problem :)
	for _, line := range input {

		found := regexPattern.FindAllString(line, len(line)+1)
		if len(found) > 0 {
			for _, mulStr := range found {
				temp := strings.Split(mulStr, " ")
				for _, mul := range performMultiplication(temp) {
					multiplesSum += mul
				}
			}
		}
	}
	fmt.Println(multiplesSum)
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
