# Go Programming Basics Study Guide

## 1. String Manipulation

### String Basics

- Strings in Go are immutable byte slices
- UTF-8 encoded by default
- Can be created using double quotes or backticks
- Length can be measured in bytes or runes

### Key Concepts

- `len()` vs `utf8.RuneCountInString()`
- String concatenation methods
- String slicing
- Unicode handling
- String conversion

### Common Operations

```go
// String length
str := "Hello, 世界"
byteLen := len(str)        // 13
runeLen := utf8.RuneCountInString(str)  // 8

// String concatenation
s1 := "Hello"
s2 := "World"
result := s1 + " " + s2  // Using +
result = fmt.Sprintf("%s %s", s1, s2)  // Using fmt.Sprintf
result = strings.Join([]string{s1, s2}, " ")  // Using strings.Join

// String slicing
str = "Hello, World"
sub := str[0:5]  // "Hello"

// Unicode handling
for _, r := range str {
    fmt.Printf("%c", r)  // Print each rune
}
```

## 2. Number Operations

### Numeric Types

- Integers: `int`, `int8`, `int16`, `int32`, `int64`
- Unsigned: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- Floating-point: `float32`, `float64`
- Complex: `complex64`, `complex128`

### Key Concepts

- Type conversion
- Arithmetic operations
- Bitwise operations
- Overflow handling
- Precision considerations

### Common Operations

```go
// Type conversion
var i int = 42
f := float64(i)
s := strconv.Itoa(i)

// Arithmetic
a, b := 10, 3
sum := a + b
diff := a - b
prod := a * b
quot := a / b
rem := a % b

// Bitwise
x := 5
y := 3
and := x & y
or := x | y
xor := x ^ y
not := ^x
shift := x << 1
```

## 3. Control Structures

### Loops

- `for` loop (only loop in Go)
- `range` for iteration
- `break` and `continue`

### Conditionals

- `if` statement
- `if` with initialization
- `switch` statement
- `select` for channels

### Common Patterns

```go
// For loop
for i := 0; i < 10; i++ {
    // ...
}

// Range loop
for index, value := range slice {
    // ...
}

// If with initialization
if value, err := someFunction(); err != nil {
    // Handle error
}

// Switch
switch value {
case 1:
    // ...
case 2:
    // ...
default:
    // ...
}
```

## 4. Functions

### Function Basics

- Function declaration
- Multiple return values
- Named return values
- Variadic functions
- Anonymous functions

### Key Concepts

- Function types
- Closures
- Defer
- Panic and recover
- Method receivers

### Common Patterns

```go
// Basic function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Anonymous function
func() {
    fmt.Println("Anonymous function")
}()
```

## 5. Error Handling

### Error Basics

- Error interface
- Error creation
- Error wrapping
- Error checking
- Custom errors

### Key Concepts

- Error types
- Error chains
- Error handling patterns
- Panic vs error
- Error context

### Common Patterns

```go
// Custom error
type ValidationError struct {
    Field string
    Issue string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Issue)
}

// Error wrapping
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}

// Error checking
if errors.Is(err, io.EOF) {
    // Handle EOF
}

// Error type assertion
if validationErr, ok := err.(*ValidationError); ok {
    // Handle validation error
}
```

## Practice Tips

1. Start with simple string operations
2. Practice number conversions and operations
3. Master control structures
4. Understand function types and closures
5. Learn error handling patterns
6. Write test cases for your code
7. Use Go's standard library effectively
8. Follow Go idioms and best practices

## Resources

- [Go Language Specification](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)
- [Go Standard Library](https://golang.org/pkg/)
