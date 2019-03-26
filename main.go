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

	if len(args) < lineCount {
		fmt.Println("Please provide a valid cron expression such as:")
		fmt.Println("$ crongo \"*/15 0 1,15 * 1-5 /user/bin/find\"")
	} else {
		lines := parse(args)

		fmt.Println(strings.Join(lines[:], "\n"))
	}
}

type parser struct{ strings.Builder }

func parse(args []string) []string {
	fmt.Println(args)
	result := make([]string, lineCount)

	var p parser
	var values string

	for i, arg := range args[:lineCount] {
		label := labels[i]

		switch label {
		case "minute":
			values = p.parseExpression(arg, 0, 59)
		case "hour":
			values = p.parseExpression(arg, 0, 23)
		case "day of month":
			values = p.parseExpression(arg, 1, 31)
		case "month":
			values = p.parseExpression(arg, 1, 12)
		case "day of week":
			values = p.parseExpression(arg, 0, 6)
		default:
			if len(args) > lineCount {
				values = arg + " " + args[i+1]
			} else {
				values = arg
			}
		}

		result[i] = fmt.Sprintf("%- 14s%s", label, values)
	}

	return result
}

func (p *parser) parseExpression(expression string, min, max int) string {
	anyValue := regexp.MustCompile(`^\*$`)
	stepsOfValues := regexp.MustCompile(`^\*\/(\d{1,2})$`)
	rangeOfValues := regexp.MustCompile(`^(\d{1,2})-(\d{1,2})$`)

	switch {
	case anyValue.MatchString(expression):
		return p.subRange(min, max)

	case stepsOfValues.MatchString(expression):
		step := step(stepsOfValues, expression)
		return p.steppedRange(min, max, step)

	case rangeOfValues.MatchString(expression):
		min, max := minAndMax(rangeOfValues, expression)
		return p.subRange(min, max)

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

func (p *parser) formatRange(min, max int, next func(int) int) string {
	p.Reset()
	p.WriteString(strconv.Itoa(min))
	for i := min + 1; i <= max; i++ {
		p.WriteString(" ")
		p.WriteString(strconv.Itoa(next(i)))
	}

	return p.String()
}

func (p *parser) steppedRange(min, max, step int) string {
	count := max / step // discarding remainder
	return p.formatRange(min, count+min, func(value int) int {
		return (value-min)*step + min
	})
}

func (p *parser) subRange(min, max int) string {
	return p.formatRange(min, max, func(value int) int {
		return value
	})
}
