package day7

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "7b",
	Short: "Day 7, Problem b",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput(), 5, 60))
	},
}

func b(input *util.ChallengeInput, workers, weight int) int {
	head := linkTree(input)
	wip := make([]*node, workers)

	time := 0
	working := 0
	for {
		time++

		for i := 0; i < workers; i++ {
			if wip[i] != nil {
				wip[i].workRemaining--
				if wip[i].workRemaining == 0 {
					for _, child := range wip[i].children {
						child.refcount--
						if child.refcount == 0 {
							head = append(head, child)
						}
					}

					wip[i] = nil
					working--
				}
			}
		}

		sort.Slice(head, func(i, j int) bool {
			return strings.Compare(head[i].key, head[j].key) < 0
		})

		for i := 0; i < workers; i++ {
			if wip[i] == nil && len(head) > 0 {
				wip[i] = head[0]
				head = head[1:]

				wip[i].workRemaining = weight + int(rune(wip[i].key[0])-'A'+1)
				working++
			}
		}

		if len(head) == 0 && working == 0 {
			break
		}
	}

	return time - 1
}
