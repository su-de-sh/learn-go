# Word Frequency

Implement functions to analyze word frequency in text using maps.

## Requirements

1. Create the following functions:

   - `WordFrequency(text string) map[string]int`: Count word frequencies
   - `TopWords(text string, n int) []WordCount`: Get top n most frequent words
   - `WordFrequencyByLength(text string) map[int][]string`: Group words by length
   - `WordFrequencyByPrefix(text string, prefixLen int) map[string]int`: Count words by prefix
   - `WordFrequencyWithContext(text string, word string, contextSize int) map[string]int`: Find word context

2. Each function should:
   - Handle empty text
   - Handle special characters
   - Be case-insensitive
   - Handle Unicode characters

## Example Usage

```go
text := "The quick brown fox jumps over the lazy dog. The fox is quick."

// Count word frequencies
freq := WordFrequency(text)
fmt.Println(freq) // map[the:2 quick:2 brown:1 fox:2 jumps:1 over:1 lazy:1 dog:1 is:1]

// Get top 3 most frequent words
top := TopWords(text, 3)
fmt.Println(top) // [{the 2} {quick 2} {fox 2}]

// Group words by length
byLength := WordFrequencyByLength(text)
fmt.Println(byLength) // map[3:[the fox dog] 4:[lazy] 5:[quick brown jumps] 6:[over]]

// Count words by prefix
byPrefix := WordFrequencyByPrefix(text, 2)
fmt.Println(byPrefix) // map[th:2 qu:2 br:1 fo:2 ju:1 ov:1 la:1 do:1 is:1]

// Find word context
context := WordFrequencyWithContext(text, "fox", 1)
fmt.Println(context) // map[brown:1 jumps:1 is:1]
```

## Requirements

1. The functions should handle:

   - Empty text
   - Special characters
   - Unicode characters
   - Performance considerations
   - Memory usage

2. Implementation details:
   - Use efficient string operations
   - Handle UTF-8 encoding
   - Consider memory usage
   - Handle edge cases

## Tips

- Use `strings.Fields()` for word splitting
- Use `strings.ToLower()` for case-insensitive comparison
- Consider using `strings.Builder` for string manipulation
- Use `unicode.IsLetter()` for word boundary detection
- Handle punctuation appropriately
