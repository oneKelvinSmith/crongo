package main

import (
	"flag"
	"fmt"
	"strings"
)

const lineCount = 6

var labels = [lineCount]string{
	"minute",
	"hour",
	"day of month",
	"month",
	"day of week",
	"command",
}

func main() {
	flag.Parse()
	lines := parse(flag.Args())
	result := strings.Join(lines[:], "\n")
	fmt.Println(result)
}

func parse(args []string) []string {
	result := make([]string, lineCount)
	var expression string

	for i, arg := range args {
		label := labels[i]

		switch label {
		case "minute":
			expression = parseExpression(arg, 60, 1)
		case "hour":
			expression = parseExpression(arg, 24, 1)
		case "day of month":
			expression = parseExpression(arg, 31, 0)
		case "month":
			expression = parseExpression(arg, 12, 0)
		case "day of week":
			expression = parseExpression(arg, 7, 1)
		default:
			expression = arg // command
		}

		result[i] = fmt.Sprintf("%- 14s%s", label, expression)
	}

	return result
}

func parseExpression(arg string, max, offset int) string {
	fmt.Print(arg)
	switch arg {
	case "*":
		return fullRange(1, max, offset)
	default:
		fmt.Println(makeRange(1, max, offset))
		return arg
	}

}

func fullRange(min, max, offset int) string {
	return fmt.Sprintf("%d - %d", min-offset, max-offset)
}

func makeRange(min, max, step int) []int {
	result := make([]int, max)
	for i := range result {
		result[i] = i + step
	}
	return result
}
