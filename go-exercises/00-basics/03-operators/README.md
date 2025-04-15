# Operators

Practice using different operators in Go.

## Requirements

1. Create an `OperatorPractice` struct with methods:

   - `ArithmeticOperators()`: Demonstrate arithmetic operations
   - `ComparisonOperators()`: Show comparison operations
   - `LogicalOperators()`: Practice logical operations
   - `BitwiseOperators()`: Work with bitwise operations
   - `OperatorPrecedence()`: Show operator precedence

2. Each function should handle:
   - Arithmetic operations
   - Comparison operations
   - Logical operations
   - Bitwise operations
   - Operator precedence

## Example Usage

```go
// Create practice instance
practice := &OperatorPractice{}

// Arithmetic operators
practice.ArithmeticOperators()
// Output:
// Addition: 5 + 3 = 8
// Subtraction: 5 - 3 = 2
// Multiplication: 5 * 3 = 15
// Division: 5 / 3 = 1
// Modulus: 5 % 3 = 2
// Increment: 5++ = 6
// Decrement: 5-- = 4

// Comparison operators
practice.ComparisonOperators()
// Output:
// Equal: 5 == 3 = false
// Not equal: 5 != 3 = true
// Less than: 5 < 3 = false
// Less than or equal: 5 <= 3 = false
// Greater than: 5 > 3 = true
// Greater than or equal: 5 >= 3 = true

// Logical operators
practice.LogicalOperators()
// Output:
// AND: true && false = false
// OR: true || false = true
// NOT: !true = false
// XOR: true != false = true
// NAND: !(true && false) = true
// NOR: !(true || false) = false

// Bitwise operators
practice.BitwiseOperators()
// Output:
// AND: 5 & 3 = 1
// OR: 5 | 3 = 7
// XOR: 5 ^ 3 = 6
// NOT: ^5 = -6
// Left shift: 5 << 1 = 10
// Right shift: 5 >> 1 = 2

// Operator precedence
practice.OperatorPrecedence()
// Output:
// 2 + 3 * 4 = 14
// (2 + 3) * 4 = 20
// 2 + 3 * 4 / 2 = 8
// 2 + (3 * 4) / 2 = 8
```

## Requirements

1. The implementation should handle:

   - Operator precedence
   - Type compatibility
   - Overflow handling
   - Division by zero
   - Edge cases
   - Type conversion

2. Implementation details:
   - Use appropriate operators
   - Handle edge cases
   - Consider type limits
   - Use proper precedence
   - Follow Go idioms

## Tips

- Understand operator precedence
- Handle type compatibility
- Consider overflow cases
- Handle division by zero
- Use parentheses for clarity
- Follow Go idioms
- Consider performance
- Handle edge cases
- Use meaningful names
- Consider type safety
