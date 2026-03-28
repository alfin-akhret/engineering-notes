# Price Data Type

What is the correct data type to use?

Rule of thumbs:
- Always use smallest unit for backend (cents)
- Format it before sending response

Summary:
- Use int64 everywhere internally
- Store as BIGINT in DB
- Convert to float only for UI
- Never use float for calculations

So:
1. DB: BIGINT
2. Domain/Service: int64
3. Request/Response: float64

## Why?

Because numeric types like float / decimal are designed for scientific operations,
NOT for representing money.

They introduce precision errors due to how numbers are stored in binary.

For example:

**❌ Using Float**
```
priceA := 19.99
priceB := 9.99

total := priceA + priceB
fmt.Println(total)
```

Output:
```
29.979999999999997 ❌
```

**❌ Another example (multipliction)**
```
price := 79.99
qty := 3

total := price * float64(qty)
fmt.Println(total)
```
Output:
```
239.97000000000003 ❌
```

### Why this is dangerous?
In an ecommerce system:
- Cart total ≠ Order total
- Payment gateway may reject request
- Invoice mismatch
- Financial report becomes inconsistent
- Hard-to-debug rounding issues

## ✅  Correct Approach (Use Integer)
```
priceA := int64(1999) // 19.99
priceB := int64(999)  // 9.99

total := priceA + priceB
fmt.Println(total)
```

Output:
```
2998 ✅
```

### Representation

Instead of `29.98`, Store `2998`

### Then use conversion
**To cents (backend)**
```
func ToCents(amount float64) int64 {
	return int64(math.Round(amount * 100))
}
```

**To Float (Response)**
```
func ToFloat(cents int64) float64 {
	return float64(cents) / 100
}
```

## ⚠️ Important Notes
**1. Always round when converting**

Without rounding:
```
int64(79.99 * 100) // can become 7998 ❌
```

**2. Avoid mixing float and int in calculation**
- Do all calculations in int64
- Convert only at the boundary (API response)

**3. Downcasting is lossy**
Converting back:
```
7999 → 79.99
```
may introduce rounding issues again. acceptable only for display.

## Key Insight

Money is not a continuous value

It is a discrete unit (smallest denomination)

