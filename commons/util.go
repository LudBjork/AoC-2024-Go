package commons

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInput(filepath string) []string {
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
func SimpleStringToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}
