package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadInput("../inputs/problem_1.txt")
	// _, str because first element in range is the index
	for _, str := range input {
		fmt.Print(str)
	}
}

// Returns every line of input string as an array.
// Further finetuning is left to each method giving the solution
func ReadInput(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	var input []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input = append(input, line)
	}

	file.Close()
	return input
}
