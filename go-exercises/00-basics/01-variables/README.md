# Variables

Practice variable declarations and usage in Go.

## Requirements

1. Create a `VariablePractice` struct with methods:

   - `DeclareVariables()`: Demonstrate different variable declarations
   - `TypeConversion()`: Show type conversion between different types
   - `ZeroValues()`: Display zero values for different types
   - `MultipleDeclaration()`: Show multiple variable declarations
   - `Constants()`: Demonstrate constant declarations and usage

2. Each function should handle:
   - Different declaration styles
   - Type conversions
   - Zero values
   - Multiple declarations
   - Constants

## Example Usage

```go
// Create practice instance
practice := &VariablePractice{}

// Declare variables
practice.DeclareVariables()
// Output:
// var name string = "John"
// age := 30
// var (
//     firstName string
//     lastName  string
//     age       int
// )

// Type conversion
practice.TypeConversion()
// Output:
// int to float64: 42 -> 42.000000
// float64 to int: 3.14 -> 3
// string to []byte: "Hello" -> [72 101 108 108 111]

// Zero values
practice.ZeroValues()
// Output:
// string: ""
// int: 0
// float64: 0.000000
// bool: false
// []int: []

// Multiple declaration
practice.MultipleDeclaration()
// Output:
// x, y := 10, 20
// a, b, c := 1, 2, 3
// name, age, isActive := "John", 30, true

// Constants
practice.Constants()
// Output:
// Pi: 3.14159
// Max: 100
// Min: 0
// Days: 7
```

## Requirements

1. The implementation should handle:

   - Different variable types
   - Type conversions
   - Zero values
   - Multiple declarations
   - Constants
   - Type inference

2. Implementation details:
   - Use proper types
   - Handle conversions safely
   - Consider type limits
   - Use appropriate declarations
   - Follow Go idioms

## Tips

- Use appropriate types
- Handle conversions safely
- Consider type limits
- Use proper declarations
- Follow Go idioms
- Consider memory usage
- Handle edge cases
- Use meaningful names
