package solutions

import (
	"fmt"
	"ludbjork/aoc-2024/commons"
	"slices"
	"strings"
)

func SolveProblem5() {

	input := commons.ReadInput("inputs/day5.txt")
	p5_part1(input)

	p5_part2(input)

}

func p5_part1(input string) {

	orderingRuleset := getOrderingRuleset(input)
	orderingList := getPageOrderingRules(input)
	correctlyOrdered := getCorrectlyOrderedRules(orderingList, orderingRuleset)

	fmt.Println(calculateMiddleSum(correctlyOrdered))
}

func p5_part2(input string) {
	incorrectlyOrdered := getIncorrectlyOrderedRules(input)

	fmt.Println(incorrectlyOrdered)
}

func calculateMiddleSum(correctlyOrdered []string) int {
	middleSum := 0
	for i := range correctlyOrdered {
		current := strings.Split(correctlyOrdered[i], ",")
		middle := current[(len(current)+1)/2-1]
		middleSum += commons.SimpleStringToInt(middle)
	}
	return middleSum
}

func getIncorrectlyOrderedRules(input string) []string {

	orderingRules := getPageOrderingRules(input)
	ruleSet := getOrderingRuleset(input)
	correctlyOrdered := getCorrectlyOrderedRules(orderingRules, ruleSet)

	var incorrectlyOrdered []string

	for _, rule := range orderingRules {

		if !slices.Contains(correctlyOrdered, rule) {
			incorrectlyOrdered = append(incorrectlyOrdered, rule)
		}
	}

	return incorrectlyOrdered
}

func sortAllIncorrectlyOrderedRules(
	incorrectlyOrdered []string,
	ruleSet map[string][]string,
) {

}

func getCorrectlyOrderedRules(
	orderingRules []string,
	orderingRuleSet []string,
) []string {
	var correctlyOrdered []string
	for lineIndex := range orderingRules {
		line := strings.Split(orderingRules[lineIndex], ",")
		if isLineCorrectlyOrdered(line, orderingRuleSet) {
			correctlyOrdered = append(correctlyOrdered, strings.Join(line, ","))
		}

	}
	return correctlyOrdered
}

func isLineCorrectlyOrdered(line []string, orderingRuleSet []string) bool {

	for i := range line {
		after := line[i+1:]
		current := line[i]
		for j := range after {
			if compareRules(orderingRuleSet, current, after[j]) == -1 {
				return false
			}
		}
	}

	return true
}

// think of it like this "a|b" means the same as a < b
// thus if prev|next matches input return 1.
//
// if they're the same somehow return 0
func compareRules(orderingRuleSet []string, prev string, next string) int {
	var bob strings.Builder
	for _, rule := range orderingRuleSet {
		if strings.Contains(rule, prev) && strings.Contains(rule, next) {

			// check less-than
			bob.WriteString(prev)
			bob.WriteString("|")
			bob.WriteString(next)

			if strings.Contains(rule, bob.String()) {
				return 1
			}

			bob.Reset()
			bob.WriteString(next)
			bob.WriteString("|")
			bob.WriteString(prev)
			if strings.Contains(rule, bob.String()) {
				return -1
			}

			return 0
		}
	}

	return 0
}

func getOrderingRuleset(input string) []string {
	var orderRuleSet []string
	var digitHolder strings.Builder
	for pos := range input {
		if digitHolder.Len() < 5 {
			digitHolder.Write([]byte(string(input[pos])))
		}
		if digitHolder.Len() == 5 {
			if strings.Contains(digitHolder.String(), "|") {

				// store each rule individually to allow two-way comparison much
				orderRuleSet = append(orderRuleSet, digitHolder.String())
			}
			digitHolder.Reset()
		}
	}
	return orderRuleSet
}

func getPageOrderingRules(input string) []string {
	var orderingList []string
	var builder strings.Builder
	excluded := strings.Split(excludeRuleSet(input), ",")
	for i := range excluded {
		if len(excluded[i]) == 4 {
			builder.WriteString(excluded[i][:2])
			orderingList = append(orderingList, builder.String())

			builder.Reset()
			builder.WriteString(excluded[i][2:])
			builder.WriteString(",")
		} else {
			builder.WriteString(excluded[i])
			builder.WriteString(",")
		}
	}
	return orderingList
}

func excludeRuleSet(input string) string {
	orderingPart := strings.Split(input, "|")
	orderingPart = orderingPart[len(orderingPart)-1:]
	return strings.Join(orderingPart, "")[2:]
}
