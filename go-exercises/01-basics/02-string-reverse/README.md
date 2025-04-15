# String Reverse

Implement a function that reverses a string while preserving Unicode characters.

## Requirements

1. Create a function `ReverseString` that:

   - Takes a string as input
   - Returns the reversed string
   - Preserves Unicode characters (including emojis, non-ASCII characters)

2. The function should handle:
   - Empty strings
   - Single character strings
   - Strings with spaces
   - Strings with special characters
   - Strings with Unicode characters

## Example Usage

```go
result := ReverseString("Hello, 世界!")
fmt.Println(result) // "!界世 ,olleH"
```

## Requirements

1. The function should:

   - Handle UTF-8 encoded strings correctly
   - Preserve all Unicode characters
   - Work with any valid string input
   - Be efficient for large strings

2. Edge cases to handle:
   - Empty string
   - Single character
   - Palindrome strings
   - Strings with mixed ASCII and Unicode
   - Strings with combining characters

## Tips

- Use `[]rune` to handle Unicode characters
- Consider using `strings.Builder` for efficient string building
- Remember that string length in bytes != string length in runes
- Handle combining characters appropriately
