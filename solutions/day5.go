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

}

func p5_part1(input string) {

	orderingRuleset := createOrderingRuleset(input)
	orderingList := getPageOrderingList(input)
	correctlyOrdered := getCorrectlyOrderedRules(orderingList, orderingRuleset)

	fmt.Println(getMiddleSum(correctlyOrdered))
}

func getMiddleSum(correctlyOrdered []string) int {
	middleSum := 0
	for i := range correctlyOrdered {
		current := strings.Split(correctlyOrdered[i], ",")
		middle := current[(len(current)+1)/2-1]
		middleSum += commons.SimpleStringToInt(middle)
	}
	return middleSum
}

func getCorrectlyOrderedRules(
	orderingRules []string,
	orderingRuleSet map[string][]string,
) []string {
	var correctlyOrdered []string
	for lineIndex := range orderingRules {
		line := strings.Split(orderingRules[lineIndex], ",")
		isLineCorrect := true
		for i := range line {
			afterCurrent := line[i+1:]
			current := line[i]
			for j := range afterCurrent {
				if !slices.Contains(orderingRuleSet[current], afterCurrent[j]) {
					isLineCorrect = false
					break
				}
			}
		}

		if isLineCorrect {
			correctlyOrdered = append(correctlyOrdered, strings.Join(line, ","))
		}

	}
	fmt.Println(correctlyOrdered)
	return correctlyOrdered
}

func createOrderingRuleset(input string) map[string][]string {
	orderMap := make(map[string][]string)
	var digitHolder strings.Builder
	for pos := range input {
		if digitHolder.Len() < 5 {
			digitHolder.Write([]byte(string(input[pos])))
		}
		if digitHolder.Len() == 5 {
			if strings.Contains(digitHolder.String(), "|") {

				order := strings.Split(digitHolder.String(), "|")
				// store as array since one ordering could match multiple
				orderMap[order[0]] = append(orderMap[order[0]], order[1])
			}
			digitHolder.Reset()
		}
	}
	return orderMap
}

func excludeRuleSet(input string) string {
	orderingPart := strings.Split(input, "|")
	orderingPart = orderingPart[len(orderingPart)-1:]
	return strings.Join(orderingPart, "")[2:]
}

func getPageOrderingList(input string) []string {
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
