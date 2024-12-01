package main

import (
	"fmt"
	"io"
	"slices"
)

func main() {
	leftList := []int{}
	rightList := []int{}
	var left, right int

	for readIntPair(&left, &right) {
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	distance := calculateDistance(leftList, rightList)
	similarity := calculateSimilarity(leftList, rightList)

	fmt.Println("Advent of Code 2024 Day 1:")
	fmt.Printf("    Part 1 = %d\n", distance)
	fmt.Printf("    Part 2 = %d\n", similarity)
}

func readIntPair(left *int, right *int) bool {
	n, err := fmt.Scanf("%d %d\n", left, right)
	if err == io.EOF {
		return false
	}
	if n != 2 || err != nil {
		panic(err)
	}
	return true
}

func calculateDistance(leftList, rightList []int) int {
	var left, right, distance int

	for i := 0; i < len(leftList); i++ {
		left = leftList[i]
		right = rightList[i]
		if left > right {
			distance += (left - right)
		} else {
			distance += (right - left)
		}
	}

	return distance
}

func calculateSimilarity(leftList, rightList []int) int {
	rightCounts := map[int]int{}
	for _, right := range rightList {
		if _, ok := rightCounts[right]; ok {
			rightCounts[right]++
		} else {
			rightCounts[right] = 1
		}
	}

	similarity := 0

	for _, left := range leftList {
		if count, ok := rightCounts[left]; ok {
			similarity += (left * count)
		}
	}

	return similarity
}
