package day2

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2018/util"
)

var B = &cobra.Command{
	Use:   "2b",
	Short: "Day 2, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) string {
	var boxes []string
	for box := range input.Lines() {
		boxes = append(boxes, box)
	}

	for {
		candidate := boxes[0]
		boxes = boxes[1:]
		for _, box := range boxes {
			for i, c := range candidate {
				cLeft := candidate[:i]
				bLeft := box[:i]

				cRight := ""
				bRight := ""
				if i < len(candidate) {
					cRight = candidate[i+1:]
					bRight = box[i+1:]
				}

				if c != rune(box[i]) && cLeft == bLeft && cRight == bRight {
					return cLeft + cRight
				}
			}
		}
	}
}
