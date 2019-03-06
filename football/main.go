package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	log.SetOutput(os.Stdout)

	// there's definitely a more efficient way to do this.
	rawData, err := ioutil.ReadFile("../football.dat")
	if err != nil {
		log.Panicf("Failed to read file: %s", err)
	}

	//                                  Team    P     W     L     D     F           A
	re := regexp.MustCompile(`^\s+\d+\. (\w+)\s+\d+\s+\d+\s+\d+\s+\d+\s+(\d+)\s+-\s+(\d+) .*`)

	bestDelta := math.MaxInt32
	bestDayNum := ""

	for i, row := range strings.Split(string(rawData), "\n") {
		res := re.FindStringSubmatch(row)
		if res == nil {
			log.Printf("row %d: no match", i)
			continue
		}
		log.Printf("row %d: have %s %s %s", i, res[1], res[2], res[3])

		teamName := res[1]
		goalsFor := AtoiOrPanic(res[2])
		goalsAgainst := AtoiOrPanic(res[3])

		if delta := Abs(goalsFor - goalsAgainst); delta < bestDelta {
			log.Printf("row %d: %d is better than %d (team %s)\n", i, delta, bestDelta, teamName)
			bestDayNum = teamName
			bestDelta = delta
		} else {
			log.Printf("row %d: %d is not better than %d (team %s)\n", i, delta, bestDelta, teamName)
		}
	}

	log.Printf("the best team was %d with delta %d\n", bestDayNum, bestDelta)
	fmt.Println(bestDayNum)
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
