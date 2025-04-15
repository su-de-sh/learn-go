# Rate Limiter

Implement a rate limiter pattern in Go.

## Requirements

1. Create a `RateLimiter` struct with methods:

   - `NewRateLimiter(rate time.Duration) *RateLimiter`: Create new rate limiter
   - `Allow() bool`: Check if request is allowed
   - `Wait() error`: Wait for next available slot
   - `WaitWithTimeout(timeout time.Duration) error`: Wait with timeout
   - `WaitWithContext(ctx context.Context) error`: Wait with context
   - `Close()`: Close rate limiter

2. Each function should handle:
   - Rate limiting
   - Token bucket algorithm
   - Sliding window
   - Burst handling
   - Resource cleanup
   - Error conditions

## Example Usage

```go
// Create rate limiter (10 requests per second)
limiter := NewRateLimiter(time.Second / 10)

// Check if request is allowed
if limiter.Allow() {
    // Process request
}

// Wait for next available slot
err := limiter.Wait()
if err != nil {
    // Handle error
}

// Wait with timeout
err = limiter.WaitWithTimeout(2 * time.Second)
if err != nil {
    // Handle timeout
}

// Wait with context
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

err = limiter.WaitWithContext(ctx)
if err != nil {
    // Handle context cancellation
}

// Close limiter
limiter.Close()
```

## Requirements

1. The implementation should handle:

   - Rate limiting
   - Burst handling
   - Resource cleanup
   - Error conditions
   - Timeout handling
   - Context cancellation
   - Thread safety

2. Implementation details:
   - Use proper synchronization
   - Handle edge cases
   - Consider performance
   - Implement cleanup
   - Handle errors properly
   - Manage resources efficiently

## Tips

- Use token bucket algorithm
- Implement sliding window
- Handle burst requests
- Use channels for signaling
- Handle worker lifecycle
- Implement proper cleanup
- Handle memory management
- Consider performance
- Handle edge cases
- Follow Go idioms
- Use appropriate types
- Consider resource limits
