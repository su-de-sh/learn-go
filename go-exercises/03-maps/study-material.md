# Go Maps Study Guide

## 1. Map Basics

### What are Maps?

- Key-value pairs
- Unordered collection
- Dynamic size
- Reference type

### Key Concepts

- Map declaration
- Map initialization
- Map operations
- Map iteration
- Map deletion

### Common Operations

```go
// Map declaration
var m1 map[string]int
m2 := map[string]int{}
m3 := make(map[string]int)

// Map initialization
m4 := map[string]int{
    "one": 1,
    "two": 2,
}

// Map operations
m4["three"] = 3        // Add/update
delete(m4, "one")      // Delete
value := m4["two"]     // Get value
exists := m4["four"]   // Check existence
```

## 2. Map Operations

### Adding and Updating

- Direct assignment
- Multiple updates
- Conditional updates
- Batch updates

### Retrieving Values

- Direct access
- Safe access
- Default values
- Existence check

### Common Patterns

```go
// Safe value retrieval
if value, exists := m["key"]; exists {
    // Use value
}

// Conditional update
if _, exists := m["key"]; !exists {
    m["key"] = value
}

// Batch update
for k, v := range updates {
    m[k] = v
}
```

## 3. Map Iteration

### Iteration Methods

- `for range` loop
- Random order
- Concurrent iteration
- Selective iteration

### Common Patterns

```go
// Basic iteration
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Key-only iteration
for key := range m {
    // Use key
}

// Value-only iteration
for _, value := range m {
    // Use value
}

// Concurrent iteration
var mu sync.RWMutex
for key, value := range m {
    mu.RLock()
    // Use key and value
    mu.RUnlock()
}
```

## 4. Map Performance

### Memory Management

- Initial capacity
- Growth strategy
- Memory efficiency
- Garbage collection

### Optimization Techniques

- Pre-allocate capacity
- Reuse maps
- Avoid unnecessary allocations
- Use appropriate types

### Common Patterns

```go
// Pre-allocate capacity
m := make(map[string]int, 1000)

// Efficient updates
if len(m) == cap(m) {
    // Grow map
    newMap := make(map[string]int, 2*len(m))
    for k, v := range m {
        newMap[k] = v
    }
    m = newMap
}

// Memory efficient
type compactMap struct {
    keys   []string
    values []int
}
```

## 5. Map Best Practices

### Common Pitfalls

- Nil maps
- Concurrent access
- Memory leaks
- Key types

### Best Practices

- Initialize properly
- Handle nil cases
- Use appropriate types
- Consider thread safety
- Avoid unnecessary copies

### Common Patterns

```go
// Proper initialization
var m1 map[string]int        // nil map
m2 := map[string]int{}       // empty map
m3 := make(map[string]int)   // empty map with default capacity

// Thread safety
var mu sync.RWMutex
func safeSet(m map[string]int, k string, v int) {
    mu.Lock()
    defer mu.Unlock()
    m[k] = v
}

// Safe access
func safeGet(m map[string]int, k string) (int, bool) {
    mu.RLock()
    defer mu.RUnlock()
    v, ok := m[k]
    return v, ok
}
```

## 6. Advanced Map Operations

### Map Operations

- Union
- Intersection
- Difference
- Subset check

### Custom Types

- Custom key types
- Custom value types
- Nested maps
- Map of maps

### Common Patterns

```go
// Map union
func union(m1, m2 map[string]int) map[string]int {
    result := make(map[string]int)
    for k, v := range m1 {
        result[k] = v
    }
    for k, v := range m2 {
        result[k] = v
    }
    return result
}

// Map intersection
func intersection(m1, m2 map[string]int) map[string]int {
    result := make(map[string]int)
    for k, v := range m1 {
        if _, exists := m2[k]; exists {
            result[k] = v
        }
    }
    return result
}

// Nested maps
type nestedMap struct {
    data map[string]map[string]int
}

func (nm *nestedMap) get(key1, key2 string) (int, bool) {
    if m2, ok := nm.data[key1]; ok {
        return m2[key2]
    }
    return 0, false
}
```

## Practice Tips

1. Understand map internals
2. Practice map operations
3. Learn performance optimization
4. Master map patterns
5. Handle edge cases
6. Write efficient code
7. Consider thread safety
8. Follow Go idioms

## Resources

- [Go Maps in Action](https://blog.golang.org/maps)
- [Go Maps](https://golang.org/ref/spec#Map_types)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Standard Library](https://golang.org/pkg/)
