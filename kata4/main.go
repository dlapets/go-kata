package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// change to something else if you don't want to see debug output:
	log.SetOutput(os.Stdout)

	fileName, re := Mode()

	// there's definitely a more efficient way to do this.
	rawData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed to read file: %s", err)
	}

	bestName := ""
	bestDelta := 0
	examined := 0

	for i, row := range strings.Split(string(rawData), "\n") {
		res := re.FindStringSubmatch(row)
		if res == nil {
			log.Printf("row %d: no match", i)
			continue
		}
		log.Printf("row %d: have %s %s %s", i, res[1], res[2], res[3])

		name := res[1]
		lhs := AtoiOrPanic(res[2])
		rhs := AtoiOrPanic(res[3])

		if delta := Abs(lhs - rhs); delta < bestDelta || examined == 0 {
			log.Printf("row %d: %d is better than %d (name %s)\n", i, delta, bestDelta, name)
			bestName = name
			bestDelta = delta
		} else {
			log.Printf("row %d: %d is not better than %d (name %s)\n", i, delta, bestDelta, name)
		}
		examined++
	}

	if examined == 0 {
		log.Panicf("no data found")
	}

	log.Printf("the best name was %s with delta %d\n", bestName, bestDelta)
	fmt.Println(bestName)
}

func AtoiOrPanic(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("failed to convert %s to int: %s", s, err)
	}
	return i
}

func Abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func Mode() (fileName string, re *regexp.Regexp) {
	mode := "weather"
	if len(os.Args) > 1 && os.Args[1] == "football" {
		mode = "football"
	}

	switch mode {
	case "football":
		re = regexp.MustCompile(`^\s+\d+\. (\w+)\s+\d+\s+\d+\s+\d+\s+\d+\s+(\d+)\s+-\s+(\d+) .*`)
		fileName = "football.dat"
	default:
		re = regexp.MustCompile(`^\s+(\d+)\s+(\d+)\*?\s+(\d+)\*?\s`)
		fileName = "weather.dat"
	}

	log.Printf("app running in %s mode using input from %s", mode, fileName)
	return
}
