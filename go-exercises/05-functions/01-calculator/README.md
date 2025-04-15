# Calculator

Create a calculator package that implements basic arithmetic operations with error handling.

## Requirements

1. Create the following functions:

   - `Add(a, b float64) (float64, error)`
   - `Subtract(a, b float64) (float64, error)`
   - `Multiply(a, b float64) (float64, error)`
   - `Divide(a, b float64) (float64, error)`
   - `Power(base, exponent float64) (float64, error)`

2. Each function should:

   - Accept two float64 parameters
   - Return a float64 result and an error
   - Handle edge cases appropriately (e.g., division by zero)
   - Return meaningful error messages

3. Create a `Calculator` struct with methods that use these functions

## Example Usage

```go
calc := NewCalculator()
result, err := calc.Divide(10, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // 5.0
```

## Requirements

1. All functions should handle:

   - Division by zero
   - Invalid power operations (e.g., negative exponents for non-integer bases)
   - NaN and infinite values
   - Overflow/underflow conditions

2. The Calculator struct should:
   - Have a constructor function `NewCalculator()`
   - Implement all arithmetic operations as methods
   - Maintain a history of operations (optional)

## Tips

- Use `math.IsNaN()` and `math.IsInf()` to check for special values
- Consider using `errors.New()` for custom error messages
- Use `math.MaxFloat64` and `math.SmallestNonzeroFloat64` for overflow checks
- Remember to handle floating-point precision issues
