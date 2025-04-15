# Function Composition

Implement various function composition patterns in Go.

## Requirements

1. Create a `Composer` struct with methods:

   - `Compose(fs ...func(int) int) func(int) int`: Compose multiple functions
   - `Pipe(fs ...func(int) int) func(int) int`: Create a pipeline of functions
   - `Chain(fs ...func(int) int) func(int) int`: Chain functions with error handling
   - `ComposeWithContext(fs ...func(context.Context, int) (int, error)) func(context.Context, int) (int, error)`: Compose context-aware functions
   - `ComposeWithTimeout(fs ...func(int) int, timeout time.Duration) func(int) (int, error)`: Compose functions with timeout

2. Each function should handle:
   - Empty function lists
   - Nil functions
   - Error propagation
   - Context cancellation
   - Timeout handling

## Example Usage

```go
// Create composer
composer := &Composer{}

// Compose functions
addOne := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }
square := func(x int) int { return x * x }

composed := composer.Compose(addOne, double, square)
result := composed(5)  // ((5 + 1) * 2)^2 = 144

// Create pipeline
pipeline := composer.Pipe(addOne, double, square)
result = pipeline(5)  // ((5 + 1) * 2)^2 = 144

// Chain with error handling
chain := composer.Chain(
    func(x int) int { return x + 1 },
    func(x int) int { return x * 2 },
    func(x int) int { return x * x },
)
result, err := chain(5)  // ((5 + 1) * 2)^2 = 144

// Compose with context
ctx := context.Background()
ctxComposed := composer.ComposeWithContext(
    func(ctx context.Context, x int) (int, error) {
        select {
        case <-ctx.Done():
            return 0, ctx.Err()
        default:
            return x + 1, nil
        }
    },
    func(ctx context.Context, x int) (int, error) {
        select {
        case <-ctx.Done():
            return 0, ctx.Err()
        default:
            return x * 2, nil
        }
    },
)
result, err = ctxComposed(ctx, 5)  // (5 + 1) * 2 = 12

// Compose with timeout
timeoutComposed := composer.ComposeWithTimeout(
    addOne,
    double,
    square,
    time.Second,
)
result, err = timeoutComposed(5)  // ((5 + 1) * 2)^2 = 144
```

## Requirements

1. The implementation should handle:

   - Empty function lists
   - Nil functions
   - Error propagation
   - Context cancellation
   - Timeout handling
   - Function composition order

2. Implementation details:
   - Use proper error handling
   - Handle edge cases
   - Consider performance
   - Implement thread safety
   - Handle context properly

## Tips

- Use closures for composition
- Handle errors appropriately
- Consider function order
- Use context for cancellation
- Implement proper timeout handling
- Consider memory usage
- Handle edge cases
- Follow Go idioms
