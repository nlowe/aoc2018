package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nlowe/aoc2018/day1"
	"github.com/nlowe/aoc2018/day2"
)

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
	},
}

func main() {
	rootCmd.AddCommand(
		day1.A, day1.B,
		day2.A, day2.B,
	)

	flags := rootCmd.PersistentFlags()
	flags.StringP("input", "i", "", "Input File to read")

	viper.BindPFlags(flags)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
