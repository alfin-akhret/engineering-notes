# Two Sum

**Problem url**: https://leetcode.com/problems/two-sum/

	•	Category: Array / Hash Map / Two-Pointer / Pair Sum
	•	Pattern: “Find a pair of elements that satisfy a condition”
	•	Constraint: Only one solution exists, unsorted array

### 1. Initial approache: Brute Force

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

**Pattern Learned**: 
Nested loop / all combinations


### 1. Optimized Approach: one pass hash map

Algorithmn:
- traverse array once
- for each element, calculate complement = target - element
- check if complement exists in map
- if exists -> return indices
- else -> add current element to map


Time complexity: o(n)

Space complexity: o(n)

Pros: faster than brute force, only using 1 loop

Cons: need more space -> o(n), (space for speed tradeoff)

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

**Pattern Learned**:
- Hash Map Lookup (existence & index retrieval)
- Complement pattern
- One-pass / early exit optimization
- Trade-off: memory for speed

###  Key Concepts / DSA Patterns
1.	Brute Force / Nested Loop → Check all pairs
2.	Hash Map (Dictionary) → Store seen elements + index for O(1) lookup
3.	Complement Concept → target - current
4.	Early Exit / One-Pass → Stop traversal immediately after finding answer
5.	Space vs Time Trade-Off → O(n) space for O(n) time
6.	Edge Cases → duplicates, negative numbers, empty array