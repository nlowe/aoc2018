package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "8a",
	Short: "Day 8, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

type node struct {
	children []*node
	metadata []int
}

type position struct {
	wip      *node
	children int
	metadata int
}

type traversalTracker struct {
	n *node
	i int
}

func makeTree(input *util.ChallengeInput) *node {
	tokens := strings.Split(<-input.Lines(), " ")
	var stack []*position
	var next *position

	var head *node
	for _, token := range tokens {
		var i int
		var err error
		if i, err = strconv.Atoi(token); err != nil {
			panic(err)
		}

		if next == nil {
			next = &position{wip: &node{
				children: []*node{},
				metadata: []int{},
			}, children: i, metadata: -1}

			if head == nil {
				head = next.wip
			}
		} else if next.metadata == -1 {
			next.metadata = i

			if next.children != 0 {
				stack = append(stack, next)
				next = nil
			}
		} else if next.metadata > 0 {
			next.metadata--
			next.wip.metadata = append(next.wip.metadata, i)

			if next.metadata == 0 {
				wip := next

				l := len(stack)

				if l > 0 {
					stack, next = stack[:l-1], stack[l-1]
					next.wip.children = append(next.wip.children, wip.wip)
					next.children--

					if next.children > 0 {
						stack = append(stack, next)
						next = nil
					}
				}
			}
		}
	}

	return head
}

func traverse(head *node, output chan<- *node) {
	var stack []*traversalTracker
	var p *traversalTracker

	for {
		if p == nil {
			p = &traversalTracker{n: head}
		} else if len(p.n.children) == 0 || p.i == len(p.n.children) {
			output <- p.n

			l := len(stack)

			if l == 0 {
				break
			}

			stack, p = stack[:l-1], stack[l-1]
		} else {
			next := &traversalTracker{n: p.n.children[p.i]}
			p.i++
			stack = append(stack, p)
			p = next
		}
	}

	close(output)
}

func a(input *util.ChallengeInput) int {
	head := makeTree(input)

	nodes := make(chan *node)
	go traverse(head, nodes)

	sum := 0
	for node := range nodes {
		for _, i := range node.metadata {
			sum += i
		}
	}

	return sum
}
