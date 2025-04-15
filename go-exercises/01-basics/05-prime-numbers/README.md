# Prime Numbers

Implement functions to work with prime numbers using different algorithms and optimizations.

## Requirements

1. Create the following functions:

   - `IsPrime(n int) bool`: Check if a number is prime
   - `GetPrimesUpTo(n int) []int`: Get all primes up to n
   - `GetNthPrime(n int) int`: Get the nth prime number
   - `PrimeFactors(n int) []int`: Get prime factors of a number
   - `NextPrime(n int) int`: Get the next prime after n

2. Each function should:
   - Handle edge cases appropriately
   - Be efficient for large numbers
   - Handle negative numbers
   - Return appropriate errors

## Example Usage

```go
// Check if prime
result := IsPrime(17)
fmt.Println(result) // true

// Get primes up to n
primes := GetPrimesUpTo(20)
fmt.Println(primes) // [2 3 5 7 11 13 17 19]

// Get nth prime
nth := GetNthPrime(5)
fmt.Println(nth) // 11

// Get prime factors
factors := PrimeFactors(84)
fmt.Println(factors) // [2 2 3 7]

// Get next prime
next := NextPrime(17)
fmt.Println(next) // 19
```

## Requirements

1. The functions should handle:

   - Negative numbers
   - Zero and one
   - Large numbers
   - Performance considerations
   - Memory usage

2. Implementation details:
   - Use efficient algorithms (Sieve of Eratosthenes)
   - Consider using goroutines for parallel processing
   - Cache results where appropriate
   - Handle overflow cases

## Tips

- Use `uint64` for larger numbers
- Consider using `big.Int` for very large numbers
- Use the Sieve of Eratosthenes for finding multiple primes
- Consider using a bitset for memory efficiency
- Use square root optimization for primality testing
