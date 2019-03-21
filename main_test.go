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
		"minute        0 - 59",
		"hour          0 - 23",
		"day of month  1 - 31",
		"month         1 - 12",
		"day of week   0 - 6",
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
		"hour          0 - 23",
		"day of month  1 - 31",
		"month         1 - 12",
		"day of week   0 - 6",
		"command       /crongo",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}

func TestParseWithSteps(t *testing.T) {
	assert := assert.New(t)

	args := []string{"0", "9,21", "*", "*/3", "*", "/crongo"}

	lines := parse(args)

	expectedLines := []string{
		"minute        0",
		"hour          9 21",
		"day of month  1 - 31",
		"month         3 6 9 12",
		"day of week   0 - 6",
		"command       /crongo",
	}

	assert.Equal(lines, expectedLines, "could not parse cron expression")
}
