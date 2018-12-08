package day7

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "7a",
	Short: "Day 7, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", a(util.ReadInput()))
	},
}

type node struct {
	key           string
	children      []*node
	refcount      int
	workRemaining int
}

type dependency struct {
	step      string
	dependsOn string
}

func parseDependency(line string) *dependency {
	parts := strings.Split(line, " ")
	return &dependency{
		step:      parts[1],
		dependsOn: parts[7],
	}
}

func a(input *util.ChallengeInput) string {
	head := linkTree(input)

	order := ""
	for {
		sort.Slice(head, func(i, j int) bool {
			return strings.Compare(head[i].key, head[j].key) < 0
		})

		order += head[0].key
		for _, child := range head[0].children {
			child.refcount--
			if child.refcount == 0 {
				head = append(head, child)
			}
		}

		if len(head) == 1 {
			break
		}

		head = head[1:]
	}

	return order
}

func linkTree(input *util.ChallengeInput) []*node {
	nodes := map[string]*node{}
	tail := &node{children: []*node{}}
	var dependencies []*dependency

	for line := range input.Lines() {
		d := parseDependency(line)
		dependencies = append(dependencies, d)
		nodes[d.step] = &node{key: d.step, children: []*node{}}

		if _, found := nodes[d.dependsOn]; !found {
			tail.key = d.dependsOn
		}
	}

	nodes[tail.key] = tail
	var head []*node
	for _, dependency := range dependencies {
		target := nodes[dependency.step]
		target.children = append(target.children, nodes[dependency.dependsOn])
		nodes[dependency.dependsOn].refcount++
	}

	for _, node := range nodes {
		if node.refcount == 0 {
			head = append(head, node)
		}
	}

	return head
}
