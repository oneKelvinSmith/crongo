package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
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
	anyValue := regexp.MustCompile(`^\*$`)
	stepsOfValues := regexp.MustCompile(`^\*\/(\d{1,2})$`)

	switch {
	case anyValue.MatchString(expression):
		return fullRange(1, max, offset)
	case stepsOfValues.MatchString(expression):
		submatch := stepsOfValues.FindStringSubmatch(expression)
		step, err := strconv.Atoi(submatch[1])
		if err != nil {
			panic("could not parse integer after regexp match")
		}

		return steppedRange(1, max, step)
	default:
		return strings.Replace(expression, ",", " ", max-1)
	}
}

func fullRange(min, max, offset int) string {
	return fmt.Sprintf("%d - %d", min-offset, max-offset)
}

func steppedRange(min, max, step int) string {
	count := max / step // discarding remainder
	var values strings.Builder

	values.WriteString(strconv.Itoa(step))
	for i := 1; i < count; i++ {
		value := i*step + step
		values.WriteString(" ")
		values.WriteString(strconv.Itoa(value))
	}

	return values.String()
}
