## Two Sum

**Problem url**: https://leetcode.com/problems/two-sum/

**1. Initial approache: Brute Force**

Algorithmn: Nested loop, check evey pair

Time complexity: o(n^2)

Space complexity: o(1)

Pros: simple, always correct

Cons: slow for large arrays

Implementation (golang):
```
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
```
---

**1. Optimized Approach: One-Pass hash map**

Algorithmn:
- traverse array once
- for each element, calculate complement = target - element
- check if complement exists in map
- if exists -> return indices
- else -> add current element to map

Time complexity: o(n)
Space complexity: o(n)
Pros: faster than brute force, only using 1 loop
Cons: slightly complex code

Implementation example using (golang):
```
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
```