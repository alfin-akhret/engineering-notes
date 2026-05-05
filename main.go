package main

import (
	"fmt"
)

func main() {

	fmt.Println(romanToInt("MC"))
}

func romanToInt(s string) int {
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	total := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			total -= romanMap[s[i]]
		} else {
			total += romanMap[s[i]]
		}
	}
	return total
}

func RomanToInteger(s string) int {

	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	for i := 0; i < len(s); i++ {

		curr := roman[s[i]]

		if i+1 < len(s) {
			next := roman[s[i+1]]
			if curr < next {
				result += next - curr
				i++
				continue
			}
		}

		result += curr
	}

	return result
}

func twoSum(nums []int, target int) []int {

	temp := make(map[int]int)

	for i, val := range nums {
		x := target - val
		_, ok := temp[x]
		if ok {
			return []int{i, temp[x]}
		}
		temp[val] = i
	}

	return nil
}

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
