# Types

Practice working with different types in Go.

## Requirements

1. Create a `TypePractice` struct with methods:

   - `NumericTypes()`: Demonstrate numeric type operations
   - `StringOperations()`: Show string operations and methods
   - `BooleanLogic()`: Practice boolean operations
   - `ByteAndRune()`: Work with bytes and runes
   - `TypeSize()`: Display size of different types

2. Each function should handle:
   - Numeric operations
   - String manipulation
   - Boolean logic
   - Byte and rune operations
   - Type size information

## Example Usage

```go
// Create practice instance
practice := &TypePractice{}

// Numeric types
practice.NumericTypes()
// Output:
// Integer: 42
// Float: 3.14159
// Complex: (1+2i)
// Int8 range: -128 to 127
// Uint8 range: 0 to 255

// String operations
practice.StringOperations()
// Output:
// Length: 13
// Contains: true
// Index: 7
// Split: [Hello World]
// Join: Hello, World
// Trim: Hello World
// Upper: HELLO WORLD
// Lower: hello world

// Boolean logic
practice.BooleanLogic()
// Output:
// AND: false
// OR: true
// NOT: false
// XOR: true
// NAND: true
// NOR: false

// Byte and rune
practice.ByteAndRune()
// Output:
// Byte: 65 (A)
// Rune: 19990 (世)
// UTF-8 length: 3
// Rune count: 1
// String to bytes: [72 101 108 108 111]
// Bytes to string: Hello

// Type size
practice.TypeSize()
// Output:
// int: 8 bytes
// float64: 8 bytes
// string: 16 bytes
// bool: 1 byte
// byte: 1 byte
// rune: 4 bytes
```

## Requirements

1. The implementation should handle:

   - Type limits
   - Type conversions
   - Type operations
   - Type size
   - Type safety
   - Type compatibility

2. Implementation details:
   - Use appropriate types
   - Handle conversions safely
   - Consider type limits
   - Use proper operations
   - Follow Go idioms

## Tips

- Understand type limits
- Handle conversions safely
- Consider type size
- Use appropriate operations
- Follow Go idioms
- Consider memory usage
- Handle edge cases
- Use meaningful names
- Consider type safety
- Handle overflow cases
