# Cache Implementation

Implement a thread-safe cache with different eviction policies using maps.

## Requirements

1. Create a `Cache` struct with the following methods:

   - `NewCache(capacity int) *Cache`: Create a new cache with specified capacity
   - `Set(key string, value interface{})`: Add or update a key-value pair
   - `Get(key string) (interface{}, bool)`: Get value by key
   - `Delete(key string)`: Remove a key-value pair
   - `Clear()`: Remove all key-value pairs
   - `Size() int`: Get current number of items
   - `Capacity() int`: Get maximum capacity

2. Implement different eviction policies:
   - LRU (Least Recently Used)
   - FIFO (First In First Out)
   - LFU (Least Frequently Used)
   - Random

## Example Usage

```go
// Create LRU cache
cache := NewCache(3)

// Add items
cache.Set("key1", "value1")
cache.Set("key2", "value2")
cache.Set("key3", "value3")

// Get value
value, exists := cache.Get("key1")
fmt.Println(value) // "value1"

// Add item when full (evicts least recently used)
cache.Set("key4", "value4")
_, exists = cache.Get("key1")
fmt.Println(exists) // false

// Delete item
cache.Delete("key2")

// Clear cache
cache.Clear()
```

## Requirements

1. The implementation should handle:

   - Concurrent access
   - Cache eviction
   - Memory management
   - Performance considerations
   - Thread safety

2. Implementation details:
   - Use efficient data structures
   - Implement eviction policies
   - Handle edge cases
   - Consider space complexity

## Tips

- Use `sync.RWMutex` for thread safety
- Consider using `container/list` for LRU implementation
- Use maps for O(1) lookups
- Consider using generics for type safety
- Handle nil values appropriately
