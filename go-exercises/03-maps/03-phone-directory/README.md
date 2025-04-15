# Phone Directory

Implement a phone directory system using maps.

## Requirements

1. Create a `PhoneDirectory` struct with the following methods:

   - `AddContact(name, phone string) error`: Add a new contact
   - `GetPhone(name string) (string, error)`: Get phone number by name
   - `RemoveContact(name string) error`: Remove a contact
   - `SearchByName(prefix string) []string`: Search contacts by name prefix
   - `SearchByPhone(prefix string) []string`: Search contacts by phone prefix
   - `GetAllContacts() map[string]string`: Get all contacts
   - `UpdatePhone(name, newPhone string) error`: Update phone number
   - `GetContactCount() int`: Get total number of contacts

2. Each method should:
   - Handle duplicate entries
   - Handle invalid phone numbers
   - Handle non-existent contacts
   - Be efficient for large directories

## Example Usage

```go
dir := NewPhoneDirectory()

// Add contacts
dir.AddContact("John Doe", "123-456-7890")
dir.AddContact("Jane Smith", "098-765-4321")

// Get phone number
phone, err := dir.GetPhone("John Doe")
fmt.Println(phone) // "123-456-7890"

// Search by name
results := dir.SearchByName("J")
fmt.Println(results) // ["John Doe", "Jane Smith"]

// Search by phone
results = dir.SearchByPhone("123")
fmt.Println(results) // ["John Doe"]

// Update phone
dir.UpdatePhone("John Doe", "111-222-3333")

// Get all contacts
all := dir.GetAllContacts()
fmt.Println(all) // map[John Doe:111-222-3333 Jane Smith:098-765-4321]

// Remove contact
dir.RemoveContact("Jane Smith")
```

## Requirements

1. The implementation should handle:

   - Duplicate names
   - Invalid phone numbers
   - Non-existent contacts
   - Performance considerations
   - Memory usage

2. Implementation details:
   - Use efficient data structures
   - Validate phone numbers
   - Handle edge cases
   - Consider space complexity

## Tips

- Use maps for O(1) lookups
- Consider using regular expressions for phone validation
- Use `strings.HasPrefix()` for prefix searches
- Consider using a custom type for phone numbers
- Handle concurrent access safely
