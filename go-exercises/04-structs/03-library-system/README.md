# Library System

Implement a library management system using structs and methods.

## Requirements

1. Create the following structs:

   - `Book`:

     - `ID`: Unique book identifier
     - `Title`: Book title
     - `Author`: Book author
     - `ISBN`: International Standard Book Number
     - `Status`: Book status (available/borrowed)
     - `DueDate`: Due date for borrowed books
     - `Borrower`: Current borrower information

   - `Member`:

     - `ID`: Unique member identifier
     - `Name`: Member name
     - `Email`: Member email
     - `BorrowedBooks`: Slice of borrowed book IDs
     - `MembershipStatus`: Active/inactive status

   - `Library`:
     - `Books`: Map of books by ID
     - `Members`: Map of members by ID
     - `Transactions`: Slice of transaction history

2. Implement the following methods:
   - `NewLibrary() *Library`: Create new library
   - `AddBook(book *Book) error`: Add book to library
   - `AddMember(member *Member) error`: Add member to library
   - `BorrowBook(bookID, memberID string) error`: Borrow a book
   - `ReturnBook(bookID string) error`: Return a book
   - `GetBookStatus(bookID string) (string, error)`: Get book status
   - `GetMemberBooks(memberID string) ([]*Book, error)`: Get member's borrowed books
   - `SearchBooks(query string) []*Book`: Search books by title/author
   - `GetOverdueBooks() []*Book`: Get list of overdue books

## Example Usage

```go
// Create library
library := NewLibrary()

// Add book
book := &Book{
    ID:     "B001",
    Title:  "The Go Programming Language",
    Author: "Alan A. A. Donovan",
    ISBN:   "978-0134190440",
    Status: "available",
}
library.AddBook(book)

// Add member
member := &Member{
    ID:              "M001",
    Name:            "John Doe",
    Email:           "john@example.com",
    MembershipStatus: "active",
}
library.AddMember(member)

// Borrow book
library.BorrowBook("B001", "M001")

// Get book status
status, _ := library.GetBookStatus("B001")
fmt.Println(status) // "borrowed"

// Get member's books
books, _ := library.GetMemberBooks("M001")
fmt.Println(books) // [Book{ID: "B001", Title: "The Go Programming Language", ...}]

// Search books
results := library.SearchBooks("Go")
fmt.Println(results) // [Book{ID: "B001", Title: "The Go Programming Language", ...}]

// Return book
library.ReturnBook("B001")
```

## Requirements

1. The implementation should handle:

   - Book availability
   - Member eligibility
   - Due dates
   - Overdue books
   - Search functionality
   - Error conditions

2. Implementation details:
   - Use proper error handling
   - Implement thread safety
   - Handle edge cases
   - Consider data consistency
   - Implement efficient search

## Tips

- Use `sync.Mutex` for thread safety
- Use maps for O(1) lookups
- Implement proper validation
- Handle concurrent access
- Consider using indexes for search
- Use time.Time for dates
