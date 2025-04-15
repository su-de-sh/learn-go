# Go Basics Study Guide

## 1. Variables and Constants

### Variable Declaration

- `var` keyword
- Short declaration (`:=`)
- Multiple declaration
- Type inference
- Zero values

### Common Patterns

```go
// Variable declaration
var name string
var age int = 25
var isActive bool = true

// Short declaration
name := "John"
age := 30

// Multiple declaration
var (
    firstName string
    lastName  string
    age       int
)

// Constants
const (
    Pi    = 3.14159
    Max   = 100
    Min   = 0
)
```

## 2. Basic Types

### Numeric Types

- Integers: `int`, `int8`, `int16`, `int32`, `int64`
- Unsigned: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- Floating-point: `float32`, `float64`
- Complex: `complex64`, `complex128`

### Other Types

- Boolean: `bool`
- String: `string`
- Byte: `byte` (alias for `uint8`)
- Rune: `rune` (alias for `int32`)

### Common Patterns

```go
// Numeric types
var i int = 42
var f float64 = 3.14
var c complex128 = 1 + 2i

// Boolean
var isTrue bool = true
var isFalse bool = false

// String
var str string = "Hello, World!"
var rawStr string = `Raw string
with multiple lines`

// Byte and Rune
var b byte = 'A'
var r rune = '世'
```

## 3. Operators

### Arithmetic Operators

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Modulus: `%`
- Increment: `++`
- Decrement: `--`

### Comparison Operators

- Equal: `==`
- Not equal: `!=`
- Less than: `<`
- Less than or equal: `<=`
- Greater than: `>`
- Greater than or equal: `>=`

### Logical Operators

- AND: `&&`
- OR: `||`
- NOT: `!`

### Common Patterns

```go
// Arithmetic
a := 10
b := 3
sum := a + b
diff := a - b
prod := a * b
quot := a / b
rem := a % b

// Comparison
isEqual := a == b
isLess := a < b
isGreater := a > b

// Logical
isTrue := true && false
isFalse := true || false
notTrue := !true
```

## 4. Control Flow

### If Statement

- Basic if
- If with initialization
- If-else
- If-else if-else

### For Loop

- Basic for
- For with condition
- For with initialization
- Range-based for
- Infinite for

### Switch Statement

- Basic switch
- Switch with expression
- Switch with multiple cases
- Switch with fallthrough

### Common Patterns

```go
// If statement
if x > 0 {
    fmt.Println("Positive")
} else if x < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Zero")
}

// If with initialization
if value, err := someFunction(); err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Value:", value)
}

// For loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Range-based for
for index, value := range items {
    fmt.Printf("Index: %d, Value: %v\n", index, value)
}

// Switch statement
switch value {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

## 5. Arrays and Slices

### Arrays

- Fixed size
- Zero value
- Array comparison
- Array copying
- Array iteration

### Slices

- Dynamic size
- Slice operations
- Slice capacity
- Slice length
- Slice copying

### Common Patterns

```go
// Arrays
var arr [5]int
arr := [5]int{1, 2, 3, 4, 5}
arr := [...]int{1, 2, 3, 4, 5}

// Slices
var slice []int
slice := make([]int, 5)
slice := []int{1, 2, 3, 4, 5}

// Slice operations
slice = append(slice, 6)
slice = slice[:len(slice)-1]
copy(dest, src)
```

## Practice Tips

1. Start with simple variable declarations
2. Practice type conversions
3. Master control structures
4. Learn array and slice operations
5. Understand operator precedence
6. Write clean, readable code
7. Use appropriate types
8. Follow Go idioms

## Resources

- [Go Language Specification](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)
- [Go Standard Library](https://golang.org/pkg/)
