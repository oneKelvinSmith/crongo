package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWithDefaults(t *testing.T) {
	assert := assert.New(t)

	args := []string{"*", "*", "*", "*", "*", "/crongo"}

	lines := parse(args)

	expectedLines := []string{
		"minute        0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59",
		"hour          0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23",
		"day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
		"month         1 2 3 4 5 6 7 8 9 10 11 12",
		"day of week   0 1 2 3 4 5 6",
		"command       /crongo",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}

func TestParseWithSimpleValues(t *testing.T) {
	assert := assert.New(t)

	args := []string{"0", "21", "29", "1", "5", "/crongo"}

	lines := parse(args)

	expectedLines := []string{
		"minute        0",
		"hour          21",
		"day of month  29",
		"month         1",
		"day of week   5",
		"command       /crongo",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}

func TestParseWithListOfValues(t *testing.T) {
	assert := assert.New(t)

	args := []string{"5,7,9", "*", "*", "*", "*", "/crongo"}

	lines := parse(args)

	expectedLines := []string{
		"minute        5 7 9",
		"hour          0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23",
		"day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
		"month         1 2 3 4 5 6 7 8 9 10 11 12",
		"day of week   0 1 2 3 4 5 6",
		"command       /crongo",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}

func TestParseWithSteps(t *testing.T) {
	assert := assert.New(t)

	args := []string{"0", "*/12", "*/4", "*/3", "*", "/crongo"}

	lines := parse(args)

	expectedLines := []string{
		"minute        0",
		"hour          0 12",
		"day of month  1 5 9 13 17 21 25 29",
		"month         1 4 7 10 13",
		"day of week   0 1 2 3 4 5 6",
		"command       /crongo",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}

func TestParseWithRange(t *testing.T) {
	assert := assert.New(t)

	args := []string{"42-49", "*", "*", "8-9", "3-5", "/crongo"}

	lines := parse(args)

	expectedLines := []string{
		"minute        42 43 44 45 46 47 48 49",
		"hour          0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23",
		"day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
		"month         8 9",
		"day of week   3 4 5",
		"command       /crongo",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}

func TestParseWithComplexInput(t *testing.T) {
	assert := assert.New(t)

	args := []string{"*/15", "0", "1,15", "*", "1-5", "/usr/bin/find"}

	lines := parse(args)

	expectedLines := []string{
		"minute        0 15 30 45",
		"hour          0",
		"day of month  1 15",
		"month         1 2 3 4 5 6 7 8 9 10 11 12",
		"day of week   1 2 3 4 5",
		"command       /usr/bin/find",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}
