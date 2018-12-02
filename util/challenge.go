package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type ChallengeInput struct {
	scanner *bufio.Scanner

	lines chan string
}

func TestInput(input string) *ChallengeInput {
	return newInputFromReader(strings.NewReader(input), nil)
}

func ReadInput() *ChallengeInput {
	if len(os.Args) < 2 {
		panic("No input")
	}

	var err error
	var f *os.File
	if f, err = os.Open(os.Args[1]); err != nil {
		panic(err)
	} else {
		return newInputFromReader(f, f)
	}
}

func newInputFromReader(r io.Reader, c io.Closer) *ChallengeInput {
	result := &ChallengeInput{}

	result.scanner = bufio.NewScanner(r)

	result.lines = make(chan string)

	go func() {
		defer func() {
			if c != nil {
				c.Close()
			}
		}()

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
