# Go Slices Study Guide

## 1. Slice Basics

### What are Slices?

- Dynamic arrays in Go
- Built on top of arrays
- Three components:
  - Pointer to underlying array
  - Length (len)
  - Capacity (cap)

### Key Concepts

- Slice declaration
- Slice initialization
- Slice capacity
- Slice length
- Slice bounds

### Common Operations

```go
// Slice declaration
var s1 []int
s2 := []int{1, 2, 3}
s3 := make([]int, 5, 10)

// Slice from array
arr := [5]int{1, 2, 3, 4, 5}
s4 := arr[1:3]  // [2, 3]

// Slice operations
len := len(s2)    // 3
cap := cap(s2)    // 3
```

## 2. Slice Manipulation

### Adding Elements

- `append` function
- Capacity growth
- Multiple append
- Append to nil slice

### Removing Elements

- Remove from beginning
- Remove from end
- Remove from middle
- Remove multiple elements

### Common Patterns

```go
// Adding elements
s := []int{1, 2, 3}
s = append(s, 4)           // [1, 2, 3, 4]
s = append(s, 5, 6)       // [1, 2, 3, 4, 5, 6]
s = append(s, []int{7, 8}...)  // [1, 2, 3, 4, 5, 6, 7, 8]

// Removing elements
s = s[1:]    // Remove first element
s = s[:len(s)-1]  // Remove last element
s = append(s[:2], s[3:]...)  // Remove element at index 2
```

## 3. Slice Operations

### Copying Slices

- `copy` function
- Deep copy
- Partial copy
- Copy to empty slice

### Slicing

- Sub-slices
- Full slice expression
- Slice bounds
- Slice capacity

### Common Patterns

```go
// Copying slices
src := []int{1, 2, 3, 4, 5}
dst := make([]int, len(src))
copy(dst, src)

// Sub-slicing
s := []int{1, 2, 3, 4, 5}
sub := s[1:3]     // [2, 3]
sub = s[1:3:4]    // [2, 3] with capacity 3
```

## 4. Slice Performance

### Memory Management

- Capacity growth
- Memory allocation
- Garbage collection
- Memory efficiency

### Optimization Techniques

- Pre-allocate capacity
- Reuse slices
- Avoid unnecessary allocations
- Use appropriate initial capacity

### Common Patterns

```go
// Pre-allocate capacity
s := make([]int, 0, 1000)  // Pre-allocate for 1000 elements

// Reuse slice
s = s[:0]  // Clear slice while keeping capacity

// Efficient append
if len(s) == cap(s) {
    // Grow slice
    newSlice := make([]int, len(s), 2*cap(s))
    copy(newSlice, s)
    s = newSlice
}
```

## 5. Slice Best Practices

### Common Pitfalls

- Nil slices vs empty slices
- Slice capacity leaks
- Concurrent access
- Slice sharing

### Best Practices

- Initialize properly
- Handle nil cases
- Use appropriate capacity
- Consider thread safety
- Avoid unnecessary copies

### Common Patterns

```go
// Proper initialization
var s1 []int        // nil slice
s2 := []int{}       // empty slice
s3 := make([]int, 0)  // empty slice with capacity 0

// Thread safety
var mu sync.Mutex
func safeAppend(s *[]int, v int) {
    mu.Lock()
    defer mu.Unlock()
    *s = append(*s, v)
}

// Avoid capacity leaks
func process(s []int) {
    s = s[:0]  // Clear slice
    // Process elements
}
```

## 6. Advanced Slice Operations

### Sorting

- Built-in sort
- Custom sorting
- Stable sorting
- Partial sorting

### Filtering

- Filter elements
- Map elements
- Reduce elements
- Find elements

### Common Patterns

```go
// Sorting
s := []int{3, 1, 4, 1, 5}
sort.Ints(s)  // [1, 1, 3, 4, 5]

// Custom sorting
sort.Slice(s, func(i, j int) bool {
    return s[i] > s[j]  // Reverse sort
})

// Filtering
func filter(s []int, predicate func(int) bool) []int {
    result := make([]int, 0, len(s))
    for _, v := range s {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}
```

## Practice Tips

1. Understand slice internals
2. Practice slice manipulation
3. Learn performance optimization
4. Master slice operations
5. Handle edge cases
6. Write efficient code
7. Consider thread safety
8. Follow Go idioms

## Resources

- [Go Slices: Usage and Internals](https://blog.golang.org/slices-intro)
- [Go Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Standard Library](https://golang.org/pkg/)
