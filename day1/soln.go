// Day1: https://adventofcode.com/2020/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func importInput(testData bool) (numbers []int) {
	if testData {
		return getTestData()
	}

	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	return numbers
}

func getTestData() (numbers []int) {
	return []int{1721, 979, 366, 299, 675, 1456}
}

func soln1(input []int) (output int) {
	for index, value := range input {
		for _, nextValue := range input[index+1:] {
			if value+nextValue == 2020 {
				return value * nextValue
			}
		}
	}

	return
}

func soln2(input []int) (output int) {
	for index, value := range input {
		for nextIndex, nextValue := range input[index+1:] {
			for _, thirdValue := range input[nextIndex+1:] {
				if value+nextValue+thirdValue == 2020 {
					return value * nextValue * thirdValue
				}
			}
		}
	}

	return
}

func main() {
	input := importInput(false)

	soln1 := soln1(input)
	soln2 := soln2(input)

	fmt.Println("Solution 1:", soln1)
	fmt.Println("Solution 2:", soln2)
}
