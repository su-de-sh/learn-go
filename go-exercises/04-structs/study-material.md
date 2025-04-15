# Go Structs Study Guide

## 1. Struct Basics

### What are Structs?

- User-defined types
- Group related data
- Custom data structures
- Value types

### Key Concepts

- Struct declaration
- Struct initialization
- Struct fields
- Struct tags
- Struct methods

### Common Operations

```go
// Struct declaration
type Person struct {
    Name    string
    Age     int
    Address string
}

// Struct initialization
p1 := Person{"John", 30, "New York"}
p2 := Person{Name: "Jane", Age: 25}
p3 := new(Person)

// Struct methods
func (p Person) String() string {
    return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}
```

## 2. Struct Methods

### Method Types

- Value receivers
- Pointer receivers
- Interface methods
- Method chaining

### Common Patterns

```go
// Value receiver
func (p Person) GetAge() int {
    return p.Age
}

// Pointer receiver
func (p *Person) SetAge(age int) {
    p.Age = age
}

// Method chaining
func (p *Person) SetName(name string) *Person {
    p.Name = name
    return p
}

// Interface implementation
type Stringer interface {
    String() string
}
```

## 3. Struct Composition

### Embedding

- Struct embedding
- Interface embedding
- Method promotion
- Field promotion

### Common Patterns

```go
// Struct embedding
type Employee struct {
    Person
    Salary float64
}

// Interface embedding
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

// Method promotion
e := Employee{
    Person: Person{Name: "John"},
    Salary: 50000,
}
fmt.Println(e.Name)  // Access promoted field
```

## 4. Struct Performance

### Memory Layout

- Field alignment
- Padding
- Memory efficiency
- Cache optimization

### Optimization Techniques

- Field ordering
- Struct size
- Memory allocation
- Zero values

### Common Patterns

```go
// Optimized struct layout
type Optimized struct {
    // 8-byte fields first
    ID        int64
    Timestamp int64
    // 4-byte fields
    Age       int32
    Status    int32
    // 1-byte fields
    Flag      bool
    Active    bool
}

// Zero value optimization
type Config struct {
    // Use zero values as defaults
    Timeout time.Duration
    Retries int
}
```

## 5. Struct Best Practices

### Common Pitfalls

- Nil pointers
- Zero values
- Field visibility
- Method receivers

### Best Practices

- Use appropriate receivers
- Handle nil cases
- Export fields properly
- Consider immutability
- Use constructor functions

### Common Patterns

```go
// Constructor function
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

// Immutable struct
type ImmutablePerson struct {
    name    string
    age     int
    address string
}

// Safe access
func (p *Person) GetName() string {
    if p == nil {
        return ""
    }
    return p.Name
}
```

## 6. Advanced Struct Operations

### Custom Types

- Type aliases
- Custom types
- Type conversion
- Type assertions

### Validation

- Field validation
- Constraint checking
- Error handling
- Custom validation

### Common Patterns

```go
// Custom type
type Age int

func (a Age) IsAdult() bool {
    return a >= 18
}

// Field validation
type User struct {
    Name  string
    Email string
    Age   int
}

func (u *User) Validate() error {
    if u.Name == "" {
        return errors.New("name is required")
    }
    if !strings.Contains(u.Email, "@") {
        return errors.New("invalid email")
    }
    if u.Age < 0 {
        return errors.New("age must be positive")
    }
    return nil
}

// Type conversion
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
```

## Practice Tips

1. Understand struct internals
2. Practice method implementation
3. Learn composition patterns
4. Master struct operations
5. Handle edge cases
6. Write efficient code
7. Consider immutability
8. Follow Go idioms

## Resources

- [Go Structs](https://golang.org/ref/spec#Struct_types)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Standard Library](https://golang.org/pkg/)
- [Go Design Patterns](https://github.com/tmrts/go-patterns)
