# Array Difference

Implement functions to find differences between arrays in different ways.

## Requirements

1. Create the following functions:

   - `ArrayDifference(arr1, arr2 []int) []int`: Elements in arr1 not in arr2
   - `SymmetricDifference(arr1, arr2 []int) []int`: Elements in either array but not both
   - `ArrayIntersection(arr1, arr2 []int) []int`: Common elements between arrays
   - `ArrayUnion(arr1, arr2 []int) []int`: All unique elements from both arrays
   - `ArrayComplement(arr1, arr2 []int) []int`: Elements in arr2 not in arr1

2. Each function should:
   - Handle empty arrays
   - Handle arrays with duplicates
   - Handle arrays of different lengths
   - Be efficient for large arrays

## Example Usage

```go
arr1 := []int{1, 2, 3, 4, 5}
arr2 := []int{4, 5, 6, 7, 8}

// Elements in arr1 not in arr2
result := ArrayDifference(arr1, arr2)
fmt.Println(result) // [1 2 3]

// Elements in either array but not both
result = SymmetricDifference(arr1, arr2)
fmt.Println(result) // [1 2 3 6 7 8]

// Common elements
result = ArrayIntersection(arr1, arr2)
fmt.Println(result) // [4 5]

// All unique elements
result = ArrayUnion(arr1, arr2)
fmt.Println(result) // [1 2 3 4 5 6 7 8]

// Elements in arr2 not in arr1
result = ArrayComplement(arr1, arr2)
fmt.Println(result) // [6 7 8]
```

## Requirements

1. The functions should handle:

   - Empty arrays
   - Arrays with duplicates
   - Arrays of different lengths
   - Performance considerations
   - Memory usage

2. Implementation details:
   - Use efficient algorithms
   - Minimize memory usage
   - Handle edge cases
   - Consider space complexity

## Tips

- Use maps for O(1) lookups
- Consider using sets for unique elements
- Use `append` for dynamic array growth
- Consider using `sort` package for sorted results
- Handle nil arrays appropriately
