# Go Functions Study Guide

## 1. Function Basics

### Function Types

- Named functions
- Anonymous functions
- Methods
- Function literals
- Function types

### Key Concepts

- Function declaration
- Function parameters
- Return values
- Function scope
- Function visibility

### Common Patterns

```go
// Named function
func add(a, b int) int {
    return a + b
}

// Anonymous function
func() {
    fmt.Println("Anonymous function")
}()

// Function type
type Operation func(int, int) int

// Function as parameter
func apply(op Operation, a, b int) int {
    return op(a, b)
}
```

## 2. Function Composition

### Composition Techniques

- Function chaining
- Pipeline pattern
- Middleware pattern
- Decorator pattern
- Currying

### Common Patterns

```go
// Function chaining
type Builder struct {
    value string
}

func (b *Builder) Add(s string) *Builder {
    b.value += s
    return b
}

// Pipeline
func pipeline(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for v := range input {
            output <- v * 2
        }
    }()
    return output
}

// Middleware
type Handler func(http.ResponseWriter, *http.Request)

func logging(h Handler) Handler {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.URL)
        h(w, r)
    }
}
```

## 3. Closures

### Closure Basics

- Lexical scope
- Variable capture
- State preservation
- Closure lifetime
- Memory management

### Common Patterns

```go
// Counter closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Factory closure
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Stateful closure
type State struct {
    value int
}

func (s *State) update() func(int) {
    return func(newValue int) {
        s.value = newValue
    }
}
```

## 4. Function Decorators

### Decorator Patterns

- Timing decorator
- Retry decorator
- Cache decorator
- Logging decorator
- Error handling decorator

### Common Patterns

```go
// Timing decorator
func timing(f func()) func() {
    return func() {
        start := time.Now()
        f()
        duration := time.Since(start)
        fmt.Printf("Function took %v\n", duration)
    }
}

// Retry decorator
func retry(attempts int, delay time.Duration) func(f func() error) func() error {
    return func(f func() error) func() error {
        return func() error {
            var err error
            for i := 0; i < attempts; i++ {
                if err = f(); err == nil {
                    return nil
                }
                time.Sleep(delay)
            }
            return err
        }
    }
}

// Cache decorator
func cache(f func(string) string) func(string) string {
    cache := make(map[string]string)
    return func(key string) string {
        if value, exists := cache[key]; exists {
            return value
        }
        value := f(key)
        cache[key] = value
        return value
    }
}
```

## 5. Advanced Function Patterns

### Functional Programming

- Pure functions
- Immutability
- Higher-order functions
- Function composition
- Recursion

### Common Patterns

```go
// Pure function
func pureAdd(a, b int) int {
    return a + b
}

// Higher-order function
func mapInts(f func(int) int, nums []int) []int {
    result := make([]int, len(nums))
    for i, n := range nums {
        result[i] = f(n)
    }
    return result
}

// Recursive function
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

## 6. Error Handling

### Error Patterns

- Error wrapping
- Error handling
- Error recovery
- Custom errors
- Error chains

### Common Patterns

```go
// Error wrapping
func process() error {
    if err := step1(); err != nil {
        return fmt.Errorf("step1 failed: %w", err)
    }
    return nil
}

// Error recovery
func safeCall(f func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    f()
    return nil
}

// Custom error
type ValidationError struct {
    Field string
    Issue string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Issue)
}
```

## Practice Tips

1. Understand function types
2. Practice function composition
3. Master closure patterns
4. Learn decorator patterns
5. Handle errors properly
6. Write pure functions
7. Use higher-order functions
8. Follow Go idioms

## Resources

- [Go Functions](https://golang.org/ref/spec#Function_types)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Standard Library](https://golang.org/pkg/)
- [Go Design Patterns](https://github.com/tmrts/go-patterns)
