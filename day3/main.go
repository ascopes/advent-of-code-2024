package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := readInput()

	fmt.Println("Advent of Code 2024 Day 3:")
	fmt.Printf("    Part 1: %d\n", part1(input))
	fmt.Printf("    Part 2: %d\n", part2(input))
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	buff, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return string(buff)
}

func parseNumber(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return result
}

func part1(input string) int64 {
	var total int64
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, match := range pattern.FindAllString(input, -1) {
		submatch := pattern.FindStringSubmatch(match)
		left := parseNumber(submatch[1])
		right := parseNumber(submatch[2])
		total += int64(left) * int64(right)
	}
	return total
}

func part2(input string) int64 {
	var total int64
	do := true
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	for _, match := range pattern.FindAllString(input, -1) {
		switch match {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				submatch := pattern.FindStringSubmatch(match)
				left := parseNumber(submatch[1])
				right := parseNumber(submatch[2])
				total += int64(left) * int64(right)
			}
		}
	}
	return total
}
