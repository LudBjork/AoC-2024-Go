package solutions

import (
	"fmt"
	"ludbjork/aoc-2024/commons"
	"slices"
	"sort"
	"strings"
)

func SolveProblem5() {

	input := commons.ReadInput("inputs/day5.txt")
	//	p5_part1(input)

	p5_part2(input)

}

func p5_part1(input string) {

	orderingRuleset := getOrderingRuleset(input)
	orderingList := getOrderings(input)
	correctlyOrdered := getCorrectlyOrderedRules(orderingList, orderingRuleset)

	fmt.Println(calculateMiddleSum(correctlyOrdered))
}

func p5_part2(input string) {
	orderingRuleset := getOrderingRuleset(input)

	incorrectlyOrdered := getIncorrectlyOrderedRules(input)

	// sorts in-place ???
	sort.Slice(incorrectlyOrdered, func(i, j int) bool {
		return compareRules(orderingRuleset, incorrectlyOrdered[i], incorrectlyOrdered[j]) == 1
	})
	fmt.Println(calculateMiddleSum(incorrectlyOrdered))

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

	orderingRules := getOrderings(input)
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
			if compareRules(orderingRuleSet, current, after[j]) != 1 {
				return false
			}
		}
	}

	return true
}

// Interpret "a|b" as "a must come before b"
// if Rule prev|next matches prev , next return 1.
// method assumes non-contradictory rules

func compareRules(orderingRuleSet []string, prev string, next string) int {
	var bob strings.Builder
	for _, rule := range orderingRuleSet {
		// not relevant to check rule if both aren't present
		if strings.Contains(rule, prev) && strings.Contains(rule, next) {

			// check less-than
			bob.WriteString(prev)
			bob.WriteString("|")
			bob.WriteString(next)

			if strings.Contains(rule, bob.String()) {
				// prev is before next
				return 1
			}
			bob.Reset()

			bob.WriteString(next)
			bob.WriteString("|")
			bob.WriteString(prev)
			if strings.Contains(rule, bob.String()) {
				return -1
			}

		}
	}

	// do nothing
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

func getOrderings(input string) []string {
	tmp := strings.Split(input, "|")
	tmp = tmp[len(tmp)-1:]

	orderings := strings.Join(tmp, "")[2:]
	orderingSlice := strings.Split(orderings, ",")
	for i := range orderingSlice {
		if len(orderingSlice[i]) == 4 {
			var bob strings.Builder
			bob.WriteString(orderingSlice[i][:2])
			bob.WriteString("\n")
			orderingSlice[i] = strings.Replace(orderingSlice[i], orderingSlice[i][:2], bob.String(), 1)
			bob.Reset()
		}
	}
	// ensure after putting in \n at every place w. xyab that we extract out
	// each line properly
	return strings.Split(strings.Join(orderingSlice, ","), "\n")
}
