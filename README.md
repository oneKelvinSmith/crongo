# Crongo

A small tool for parsing (some/most/enough?) cron expressions

### Prequisites

Golang 1.12.1 or later. See https://golang.org/doc/install for intallation instructions.

## Installation

Fetch package with
```bash
$ go get github.com/oneKelvinSmith/crongo
```

## Usage

```bash
$ crongo "*/15 0 1,15 * 1-5 /usr/bin/find"
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

[Note]: assumes that a single argument will be passed to the binary. Other arguments are ignored.

## Contributing

First clone the repo:

```bash
$ git clone https://github.com/oneKelvinSmith/crongo
```

Ensure all dependencies are met:
```bash
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
```

Make sure the tests are running:

```bash
$ go test
PASS
ok  	github.com/oneKelvinSmith/crongo	0.022s
```

To verify the commandline client, first build it:

```bash
$ go build
```

and then run binary locally with:

```bash
$ ./crongo "* * * * * /hello"                                                                                                                                                                             00:12:21
minute        0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59
hour          0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23
day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   0 1 2 3 4 5 6
command       /hellod
```

## Known issues

Edge case and error conditions have not all been gracefully handled.