# Fibonacci Sequence

Implement functions to generate Fibonacci numbers using different approaches.

## Requirements

1. Create the following functions:

   - `FibonacciRecursive(n int) int`: Recursive implementation
   - `FibonacciIterative(n int) int`: Iterative implementation
   - `FibonacciMemoized(n int) int`: Memoized implementation
   - `FibonacciChannel(n int) []int`: Channel-based implementation

2. Each function should:
   - Handle non-negative integers
   - Return appropriate error for negative numbers
   - Be efficient for large numbers
   - Handle overflow cases

## Example Usage

```go
// Recursive
result := FibonacciRecursive(10)
fmt.Println(result) // 55

// Iterative
result = FibonacciIterative(10)
fmt.Println(result) // 55

// Memoized
result = FibonacciMemoized(10)
fmt.Println(result) // 55

// Channel-based
results := FibonacciChannel(10)
fmt.Println(results) // [0 1 1 2 3 5 8 13 21 34 55]
```

## Requirements

1. The functions should handle:

   - Negative numbers (return error)
   - Large numbers (handle overflow)
   - Zero and one as special cases
   - Performance considerations

2. Implementation details:
   - Recursive: Use tail recursion where possible
   - Iterative: Use O(n) solution
   - Memoized: Cache results for efficiency
   - Channel: Generate sequence concurrently

## Tips

- Use `uint64` for larger numbers
- Consider using `big.Int` for very large numbers
- Use `defer` for cleanup in recursive solutions
- Consider using a slice for memoization
- Use buffered channels for better performance
