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
			values = parseExpression(arg, 0, 59)
		case "hour":
			values = parseExpression(arg, 0, 23)
		case "day of month":
			values = parseExpression(arg, 1, 31)
		case "month":
			values = parseExpression(arg, 1, 12)
		case "day of week":
			values = parseExpression(arg, 0, 6)
		default:
			values = arg // command
		}

		result[i] = fmt.Sprintf("%- 14s%s", label, values)
	}

	return result
}

func parseExpression(expression string, min, max int) string {
	anyValue := regexp.MustCompile(`^\*$`)
	stepsOfValues := regexp.MustCompile(`^\*\/(\d{1,2})$`)
	rangeOfValues := regexp.MustCompile(`^(\d{1,2})-(\d{1,2})$`)

	switch {
	case anyValue.MatchString(expression):
		return subRange(min, max)
	case stepsOfValues.MatchString(expression):
		submatch := stepsOfValues.FindStringSubmatch(expression)
		step, err := strconv.Atoi(submatch[1])
		if err != nil {
			panic("could not parse integer after regexp match")
		}

		return steppedRange(min, max, step)
	case rangeOfValues.MatchString(expression):
		submatch := rangeOfValues.FindStringSubmatch(expression)
		min, err := strconv.Atoi(submatch[1])
		max, err := strconv.Atoi(submatch[2])
		if err != nil {
			panic("could not parse integer after regexp match")
		}

		return subRange(min, max)
	default:
		return strings.Replace(expression, ",", " ", max-1)
	}
}

func fullRange(min, max int) string {
	return fmt.Sprintf("%d - %d", min, max)
}

func steppedRange(min, max, step int) string {
	var values strings.Builder

	count := max / step // discarding remainder

	values.WriteString(strconv.Itoa(min))
	for i := 1; i <= count; i++ {
		value := i*step + min
		values.WriteString(" ")
		values.WriteString(strconv.Itoa(value))
	}

	return values.String()
}

func subRange(min, max int) string {
	var values strings.Builder

	values.WriteString(strconv.Itoa(min))
	for i := min + 1; i <= max; i++ {
		values.WriteString(" ")
		values.WriteString(strconv.Itoa(i))
	}

	return values.String()
}
