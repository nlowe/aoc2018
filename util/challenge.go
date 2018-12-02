package util

import (
	"bufio"
	"os"
)

type ChallengeInput struct {
	file *os.File
	scanner *bufio.Scanner

	lines chan string
}

func ReadInput() *ChallengeInput {
	if len(os.Args) < 2 {
		panic("No input")
	}

	result := &ChallengeInput{}

	var err error
	if result.file, err = os.Open(os.Args[1]); err != nil {
		panic(err)
	}

	result.scanner = bufio.NewScanner(result.file)

	result.lines = make(chan string)

	go func(){
		defer result.file.Close()

		for result.scanner.Scan() {
			result.lines <- result.scanner.Text()
		}

		close(result.lines)
	}()

	return result
}

func (c *ChallengeInput) Lines() <-chan string {
	return c.lines
}