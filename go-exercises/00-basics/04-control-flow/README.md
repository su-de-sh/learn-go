# Control Flow

Practice control flow structures in Go.

## Requirements

1. Create a `ControlFlowPractice` struct with methods:

   - `IfStatements()`: Demonstrate if statement variations
   - `ForLoops()`: Show different for loop patterns
   - `SwitchStatements()`: Practice switch statement usage
   - `BreakAndContinue()`: Work with break and continue
   - `NestedControl()`: Show nested control structures

2. Each function should handle:
   - If statement variations
   - For loop patterns
   - Switch statement cases
   - Break and continue usage
   - Nested control structures

## Example Usage

```go
// Create practice instance
practice := &ControlFlowPractice{}

// If statements
practice.IfStatements()
// Output:
// Basic if: x is positive
// If-else: x is negative
// If-else if-else: x is zero
// If with initialization: value is 42
// If with multiple conditions: x is between 0 and 10

// For loops
practice.ForLoops()
// Output:
// Basic for: 0 1 2 3 4
// For with condition: 0 1 2 3 4
// For with initialization: 0 1 2 3 4
// Range-based for: 0 1 2 3 4
// Infinite for (with break): 0 1 2 3 4

// Switch statements
practice.SwitchStatements()
// Output:
// Basic switch: One
// Switch with expression: Even
// Switch with multiple cases: Small
// Switch with fallthrough: One Two
// Switch with type assertion: String

// Break and continue
practice.BreakAndContinue()
// Output:
// With break: 0 1 2
// With continue: 0 1 3 4
// With labeled break: 0 1
// With labeled continue: 0 1 3 4

// Nested control
practice.NestedControl()
// Output:
// Nested if: x is positive and even
// Nested for: 0 1 2
// Nested switch: One and positive
// Mixed nesting: 0 1 2
```

## Requirements

1. The implementation should handle:

   - Different control structures
   - Edge cases
   - Error conditions
   - Loop termination
   - Control flow optimization
   - Code readability

2. Implementation details:
   - Use appropriate structures
   - Handle edge cases
   - Consider performance
   - Use proper nesting
   - Follow Go idioms

## Tips

- Use appropriate control structures
- Handle edge cases
- Consider performance
- Use proper nesting
- Follow Go idioms
- Consider readability
- Handle errors properly
- Use meaningful names
- Consider code organization
- Handle infinite loops
- Use break and continue wisely
- Consider switch fallthrough
- Handle type assertions
- Consider loop efficiency
