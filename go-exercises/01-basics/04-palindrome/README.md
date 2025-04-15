# Palindrome Checker

Implement functions to check if a string is a palindrome, considering different requirements.

## Requirements

1. Create the following functions:

   - `IsPalindrome(s string) bool`: Basic palindrome check
   - `IsPalindromeIgnoreCase(s string) bool`: Case-insensitive check
   - `IsPalindromeIgnoreSpaces(s string) bool`: Ignore spaces
   - `IsPalindromeIgnoreSpecial(s string) bool`: Ignore special characters
   - `IsPalindromeUnicode(s string) bool`: Handle Unicode characters

2. Each function should:
   - Handle empty strings
   - Handle single character strings
   - Handle Unicode characters
   - Be efficient for large strings

## Example Usage

```go
// Basic check
result := IsPalindrome("racecar")
fmt.Println(result) // true

// Case-insensitive
result = IsPalindromeIgnoreCase("Racecar")
fmt.Println(result) // true

// Ignore spaces
result = IsPalindromeIgnoreSpaces("race car")
fmt.Println(result) // true

// Ignore special characters
result = IsPalindromeIgnoreSpecial("race, car!")
fmt.Println(result) // true

// Unicode
result = IsPalindromeUnicode("上海海上")
fmt.Println(result) // true
```

## Requirements

1. The functions should handle:

   - Empty strings
   - Single character strings
   - Unicode characters
   - Special characters
   - Spaces and whitespace
   - Case sensitivity

2. Implementation details:
   - Use efficient string operations
   - Handle UTF-8 encoding
   - Consider memory usage
   - Handle edge cases

## Tips

- Use `strings.ToLower()` for case-insensitive checks
- Use `unicode.IsLetter()` and `unicode.IsNumber()` for special character handling
- Consider using `strings.Builder` for string manipulation
- Use `[]rune` for Unicode character handling
- Remember to handle combining characters
