# Array Merge

Implement functions to merge arrays in different ways.

## Requirements

1. Create the following functions:

   - `MergeSorted(arr1, arr2 []int) []int`: Merge two sorted arrays
   - `MergeAlternating(arr1, arr2 []int) []int`: Merge arrays alternately
   - `MergeWithPriority(arr1, arr2 []int, priority func(int) int) []int`: Merge with custom priority
   - `MergeUnique(arr1, arr2 []int) []int`: Merge arrays removing duplicates
   - `MergeInPlace(arr1, arr2 []int) []int`: Merge arrays in place

2. Each function should:
   - Handle empty arrays
   - Handle arrays of different lengths
   - Handle duplicate elements
   - Be efficient for large arrays

## Example Usage

```go
arr1 := []int{1, 3, 5, 7}
arr2 := []int{2, 4, 6, 8}

// Merge sorted arrays
result := MergeSorted(arr1, arr2)
fmt.Println(result) // [1 2 3 4 5 6 7 8]

// Merge alternately
result = MergeAlternating(arr1, arr2)
fmt.Println(result) // [1 2 3 4 5 6 7 8]

// Merge with custom priority
result = MergeWithPriority(arr1, arr2, func(x int) int { return x % 2 })
fmt.Println(result) // [2 4 6 8 1 3 5 7]

// Merge unique elements
arr1 = []int{1, 2, 3, 3, 4}
arr2 = []int{3, 4, 5, 5, 6}
result = MergeUnique(arr1, arr2)
fmt.Println(result) // [1 2 3 4 5 6]

// Merge in place
arr1Copy := make([]int, len(arr1))
copy(arr1Copy, arr1)
result = MergeInPlace(arr1Copy, arr2)
fmt.Println(result) // [1 2 3 4 5 6]
```

## Requirements

1. The functions should handle:

   - Empty arrays
   - Arrays of different lengths
   - Duplicate elements
   - Performance considerations
   - Memory usage

2. Implementation details:
   - Use efficient algorithms
   - Minimize memory usage
   - Handle edge cases
   - Consider space complexity

## Tips

- Use two-pointer technique for sorted arrays
- Consider using maps for unique elements
- Use `append` for dynamic array growth
- Consider using `sort` package for sorting
- Handle nil arrays appropriately
