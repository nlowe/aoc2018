package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nlowe/aoc2018/day13"

	"github.com/nlowe/aoc2018/day12"

	"github.com/nlowe/aoc2018/day10"
	"github.com/nlowe/aoc2018/day11"
	"github.com/nlowe/aoc2018/day9"

	"github.com/nlowe/aoc2018/day8"

	"github.com/nlowe/aoc2018/day7"

	"github.com/nlowe/aoc2018/day6"

	"github.com/nlowe/aoc2018/day5"

	"github.com/nlowe/aoc2018/day4"

	"github.com/nlowe/aoc2018/day3"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nlowe/aoc2018/day1"
	"github.com/nlowe/aoc2018/day2"
)

var start time.Time

var rootCmd = &cobra.Command{
	Use:   "aoc2018",
	Short: "Advent of Code 2018 Solutions",
	Long:  "Golang implementations for the 2018 Advent of Code problems",
	Args:  cobra.ExactArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			viper.Set("input", args[0])
		} else if !cmd.Flags().Changed("input") {
			fmt.Println("Input File Required")
			os.Exit(1)
		}
		start = time.Now()
	},
	PersistentPostRun: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Took %s\n", time.Since(start))
	},
}

func main() {
	rootCmd.AddCommand(
		day1.A, day1.B,
		day2.A, day2.B,
		day3.A, day3.B,
		day4.A, day4.B,
		day5.A, day5.B,
		day6.A, day6.B,
		day7.A, day7.B,
		day8.A, day8.B,
		day9.A, day9.B,
		day10.A,
		day11.A, day11.B,
		day12.A, day12.B,
		day13.A, day13.B,
	)

	flags := rootCmd.PersistentFlags()
	flags.StringP("input", "i", "", "Input File to read")

	viper.BindPFlags(flags)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
