# Roman to integer

**Problem url**: https://leetcode.com/problems/roman-to-integer/description/

**Problem**: Convert Roman numeral string → integer, handling subtraction rule (I before V/X, etc.)

## My 1st Approach
```
func romanToInt(s string) int {
    roman := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var (
		temp   byte
		result int
	)

	for i := 0; i < len(s); i++ {
		if temp != 0 {
			x := string(temp) + string(s[i])
			result += roman[x]
			temp = 0
			continue
		}

		if i+1 < len(s) {
			if s[i] == 'I' && (s[i+1] == 'X' || s[i+1] == 'V') {
				temp = s[i]
				continue
			}

			if s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
				temp = s[i]
				continue
			}

			if s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
				temp = s[i]
				continue
			}
		}

		result += roman[string(s[i])]
	}

	return result
}
```

***notes:***
1. long hardcoded repetitive logics
2. unnecessary string conversions
3. unnecessary string mapping. eg: `"IV": 4,` 
4. using `roman := map[string]int` instead of `roman := map[byte]int`

## My 2nd Approach
```
func romanToInt(s string) int {
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
```

***notes:*** 
1. More compact logic
2. still using `continue` and `i++`

## Optimal Solution
From other leetcode users
```
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
```

## Conclusion
**Key Data Structures:**
- map[byte]int → fast lookup for Roman symbols
- Use byte for ASCII (Roman numerals), rune only for Unicode

**Core Algorithm (Pattern-Based):**
1. Iterate string left → right
2. If current < next → subtract current
3. Else → add current

**Exmaple:**
```
if roman[s[i]] < roman[s[i+1]] {
    total -= roman[s[i]]
} else {
    total += roman[s[i]]
}
```

**Complexity:**
- Time: O(n)
- Space: O(1)

**Lessons Learned:**
- Avoid unnecessary string concatenation / extra state
- Compare adjacent characters to detect subtraction → simpler & clean
- Pattern recognition > hardcoding all cases
- Guard loop for out-of-bounds (i < len(s)-1)
- Elegant solution = mathematical transformation (+/-) instead of multiple conditional cases
