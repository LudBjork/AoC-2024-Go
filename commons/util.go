package commons

import (
	"bufio"
	"os"
	"strconv"
)

// DANGER: Depends on where program is ran from. I.e. use project root dir
// always!!!
func DeprecatedReadInput(filepath string) []string {
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

func ReadInput(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	file.Close()
	return input

}

func SimpleStringToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

func StringSliceToIntSlice(in []string) []int {
	var out []int
	for _, value := range in {
		out = append(out, SimpleStringToInt(value))
	}
	return out
}

func ComputeDistance(left int, right int) int {
	if left > right {
		return left - right
	}
	return right - left
}
