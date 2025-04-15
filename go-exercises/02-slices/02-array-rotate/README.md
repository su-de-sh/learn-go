# Array Rotation

Implement functions to rotate arrays in different ways.

## Requirements

1. Create the following functions:

   - `RotateLeft(arr []int, k int) []int`: Rotate array left by k positions
   - `RotateRight(arr []int, k int) []int`: Rotate array right by k positions
   - `RotateInPlace(arr []int, k int)`: Rotate array in place
   - `RotateByReversal(arr []int, k int)`: Rotate using reversal algorithm
   - `RotateByGCD(arr []int, k int)`: Rotate using GCD algorithm

2. Each function should:
   - Handle empty arrays
   - Handle single element arrays
   - Handle k > array length
   - Be efficient for large arrays

## Example Usage

```go
arr := []int{1, 2, 3, 4, 5}

// Rotate left
result := RotateLeft(arr, 2)
fmt.Println(result) // [3 4 5 1 2]

// Rotate right
result = RotateRight(arr, 2)
fmt.Println(result) // [4 5 1 2 3]

// Rotate in place
arrCopy := make([]int, len(arr))
copy(arrCopy, arr)
RotateInPlace(arrCopy, 2)
fmt.Println(arrCopy) // [3 4 5 1 2]

// Rotate by reversal
result = RotateByReversal(arr, 2)
fmt.Println(result) // [3 4 5 1 2]

// Rotate by GCD
result = RotateByGCD(arr, 2)
fmt.Println(result) // [3 4 5 1 2]
```

## Requirements

1. The functions should handle:

   - Empty arrays
   - Single element arrays
   - k > array length
   - Negative k values
   - Performance considerations

2. Implementation details:
   - Use efficient algorithms
   - Minimize memory usage
   - Handle edge cases
   - Consider space complexity

## Tips

- Use modulo operation for k > array length
- Consider using `append` for non-in-place rotations
- Use three-reversal algorithm for in-place rotation
- Consider using GCD for optimal rotations
- Handle negative rotations appropriately
