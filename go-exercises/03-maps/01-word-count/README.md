# Word Count

Define a function `WordCount` that accepts a string as an argument.

`WordCount` should return a map where the keys are words from the input string and the values are the number of times each word appears. Words should be case-insensitive, and punctuation should be ignored.

```go
WordCount("Hello, hello! How are you?")
// returns map[string]int{
//     "hello": 2,
//     "how": 1,
//     "are": 1,
//     "you": 1,
// }
```

## Requirements

1. The function should be named `WordCount`
2. It should accept a `string` parameter
3. It should return a `map[string]int`
4. Words should be case-insensitive
5. Punctuation should be ignored
6. The function should handle empty strings
7. The function should handle strings with only punctuation

## Tips

- Use `strings.Fields()` to split the string into words
- Use `strings.ToLower()` for case-insensitive comparison
- Use `strings.Trim()` to remove punctuation
- Consider using a regular expression to clean the words
- Remember to handle edge cases (empty string, only punctuation)
