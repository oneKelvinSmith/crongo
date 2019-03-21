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
	args := strings.Split(flag.Arg(0), " ")
	lines := parse(args)
	result := strings.Join(lines[:], "\n")
	fmt.Println(result)
}

func parse(args []string) []string {
	result := make([]string, lineCount)
	var values string

	for i, arg := range args {
		label := labels[i]

		switch label {
		case "minute":
			values = parseExpression(arg, 60, 1)
		case "hour":
			values = parseExpression(arg, 24, 1)
		case "day of month":
			values = parseExpression(arg, 31, 0)
		case "month":
			values = parseExpression(arg, 12, 0)
		case "day of week":
			values = parseExpression(arg, 7, 1)
		default:
			values = arg // command
		}

		result[i] = fmt.Sprintf("%- 14s%s", label, values)
	}

	return result
}

func parseExpression(expression string, max, offset int) string {
	switch expression {
	case "*":
		return fullRange(1, max, offset)
	default:
		return strings.Replace(expression, ",", " ", max-1)
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
