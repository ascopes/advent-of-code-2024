package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := readReports()

	fmt.Println("Advent of Code 2024 Day 2:")
	fmt.Printf("    Part 1: %d\n", part1(reports))
	fmt.Printf("    Part 2: %d\n", part2(reports))
}

func readReports() [][]int {
	reports := [][]int{}
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		values := strings.Split(strings.TrimSpace(line), " ")
		report := make([]int, len(values))
		for i, value := range values {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			report[i] = intValue
		}

		reports = append(reports, report)
	}

	return reports
}

func isSafe(report []int) bool {
	inc := true
	dec := true
	valid := true

	for i := 1; i < len(report); i++ {
		x := report[i-1]
		y := report[i]
		var diff int
		if x > y {
			diff = x - y
			dec = false
		} else {
			diff = y - x
			inc = false
		}
		if diff < 1 || diff > 3 {
			valid = false
		}
	}
	return valid && (inc || dec)
}

func part1(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func part2(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
			continue
		}
		for i := 0; i < len(report); i++ {
			newReport := append([]int{}, report[0:i]...)
			newReport = append(newReport, report[i+1:]...)
			if isSafe(newReport) {
				count++
				break
			}
		}
	}
	return count
}
