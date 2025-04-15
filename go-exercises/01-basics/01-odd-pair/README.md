# Odd Pair

Define a function `OddPair` that accepts a slice of integers as an argument.

`OddPair` should return a new slice containing the first two odd numbers from the
original slice:

```go
OddPair([]int{1, 2, 3, 4, 5}) // returns []int{1, 3}
```

If fewer than two odd numbers exist in the original slice, return a slice with the only odd number
or an empty slice:

```go
OddPair([]int{10, 15, 20}) // returns []int{15}
OddPair([]int{2, 4, 6, 8}) // returns []int{}
```

## Requirements

1. The function should be named `OddPair`
2. It should accept a `[]int` parameter
3. It should return a `[]int`
4. The function should handle nil or empty input slices
5. The function should not modify the input slice

## Tips

- Use a loop to iterate through the input slice
- Use the modulo operator `%` to check if a number is odd
- Consider using `append()` to build the result slice
- Remember to handle edge cases (nil slice, empty slice, no odd numbers)
