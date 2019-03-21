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
	var builder strings.Builder
	var values string

	for i, arg := range args {
		label := labels[i]

		switch label {
		case "minute":
			values = parseExpression(&builder, arg, 0, 59)
		case "hour":
			values = parseExpression(&builder, arg, 0, 23)
		case "day of month":
			values = parseExpression(&builder, arg, 1, 31)
		case "month":
			values = parseExpression(&builder, arg, 1, 12)
		case "day of week":
			values = parseExpression(&builder, arg, 0, 6)
		default:
			values = arg // command
		}

		result[i] = fmt.Sprintf("%- 14s%s", label, values)
	}

	return result
}

func parseExpression(b *strings.Builder, expression string, min, max int) string {
	anyValue := regexp.MustCompile(`^\*$`)
	stepsOfValues := regexp.MustCompile(`^\*\/(\d{1,2})$`)
	rangeOfValues := regexp.MustCompile(`^(\d{1,2})-(\d{1,2})$`)

	switch {
	case anyValue.MatchString(expression):
		return subRange(b, min, max)

	case stepsOfValues.MatchString(expression):
		step := step(stepsOfValues, expression)
		return steppedRange(b, min, max, step)

	case rangeOfValues.MatchString(expression):
		min, max := minAndMax(rangeOfValues, expression)
		return subRange(b, min, max)

	default:
		return strings.Replace(expression, ",", " ", max-1)
	}
}

func step(matcher *regexp.Regexp, expression string) (step int) {
	submatch := matcher.FindStringSubmatch(expression)
	step, err := strconv.Atoi(submatch[1])
	if err != nil {
		panic("could not parse integer after regexp match")
	}
	return
}

func minAndMax(matcher *regexp.Regexp, expression string) (min, max int) {
	submatch := matcher.FindStringSubmatch(expression)
	min, err := strconv.Atoi(submatch[1])
	max, err = strconv.Atoi(submatch[2])
	if err != nil {
		panic("could not parse integer after regexp match")
	}
	return
}

func formatRange(b *strings.Builder, min, max int, next func(int) int) string {
	b.Reset()
	b.WriteString(strconv.Itoa(min))
	for i := min + 1; i <= max; i++ {
		b.WriteString(" ")
		b.WriteString(strconv.Itoa(next(i)))
	}

	return b.String()
}

func steppedRange(b *strings.Builder, min, max, step int) string {
	count := max / step // discarding remainder
	return formatRange(b, min, count+min, func(value int) int {
		return (value-min)*step + min
	})
}

func subRange(b *strings.Builder, min, max int) string {
	return formatRange(b, min, max, func(value int) int {
		return value
	})
}
