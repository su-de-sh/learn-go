# Array Unique

Implement functions to handle unique elements in arrays in different ways.

## Requirements

1. Create the following functions:

   - `Unique(arr []int) []int`: Remove all duplicates
   - `UniqueInPlace(arr []int) []int`: Remove duplicates in place
   - `UniqueWithCount(arr []int) map[int]int`: Count occurrences of each element
   - `UniqueWithOrder(arr []int) []int`: Remove duplicates preserving order
   - `UniqueWithCustom(arr []int, eq func(int, int) bool) []int`: Remove duplicates with custom equality

2. Each function should:
   - Handle empty arrays
   - Handle arrays with multiple duplicates
   - Handle arrays with all unique elements
   - Be efficient for large arrays

## Example Usage

```go
arr := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

// Remove all duplicates
result := Unique(arr)
fmt.Println(result) // [1 2 3 4]

// Remove duplicates in place
arrCopy := make([]int, len(arr))
copy(arrCopy, arr)
result = UniqueInPlace(arrCopy)
fmt.Println(result) // [1 2 3 4]

// Count occurrences
counts := UniqueWithCount(arr)
fmt.Println(counts) // map[1:1 2:2 3:3 4:4]

// Remove duplicates preserving order
arr = []int{3, 1, 2, 2, 3, 1, 4}
result = UniqueWithOrder(arr)
fmt.Println(result) // [3 1 2 4]

// Remove duplicates with custom equality
result = UniqueWithCustom(arr, func(a, b int) bool { return a%2 == b%2 })
fmt.Println(result) // [3 2]
```

## Requirements

1. The functions should handle:

   - Empty arrays
   - Arrays with multiple duplicates
   - Arrays with all unique elements
   - Performance considerations
   - Memory usage

2. Implementation details:
   - Use efficient algorithms
   - Minimize memory usage
   - Handle edge cases
   - Consider space complexity

## Tips

- Use maps for O(1) lookups
- Consider using `sort` package for sorted results
- Use `append` for dynamic array growth
- Consider using `strings.Builder` for string arrays
- Handle nil arrays appropriately
