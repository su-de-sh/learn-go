# Bank Account

Implement a bank account system using structs and methods.

## Requirements

1. Create an `Account` struct with the following fields:

   - `ID`: Unique account identifier
   - `Name`: Account holder name
   - `Balance`: Current balance
   - `Transactions`: Slice of transaction history

2. Create a `Transaction` struct with:

   - `Type`: Transaction type (deposit/withdrawal)
   - `Amount`: Transaction amount
   - `Timestamp`: Transaction time
   - `Description`: Transaction description

3. Implement the following methods:
   - `NewAccount(name string) *Account`: Create new account
   - `Deposit(amount float64) error`: Add money to account
   - `Withdraw(amount float64) error`: Remove money from account
   - `Transfer(to *Account, amount float64) error`: Transfer money between accounts
   - `GetBalance() float64`: Get current balance
   - `GetTransactionHistory() []Transaction`: Get transaction history
   - `GetStatement(start, end time.Time) []Transaction`: Get statement for date range

## Example Usage

```go
// Create accounts
account1 := NewAccount("John Doe")
account2 := NewAccount("Jane Smith")

// Deposit money
account1.Deposit(1000.0)

// Withdraw money
account1.Withdraw(500.0)

// Transfer money
account1.Transfer(account2, 200.0)

// Get balance
balance := account1.GetBalance()
fmt.Println(balance) // 300.0

// Get transaction history
history := account1.GetTransactionHistory()
fmt.Println(history) // [{deposit 1000.0 2024-03-27 10:00:00} {withdrawal 500.0 2024-03-27 10:01:00} {transfer 200.0 2024-03-27 10:02:00}]

// Get statement
start := time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC)
end := time.Date(2024, 3, 27, 23, 59, 59, 0, time.UTC)
statement := account1.GetStatement(start, end)
```

## Requirements

1. The implementation should handle:

   - Negative balances
   - Invalid amounts
   - Concurrent access
   - Transaction history
   - Error conditions

2. Implementation details:
   - Use proper error handling
   - Implement thread safety
   - Handle edge cases
   - Consider data consistency

## Tips

- Use `sync.Mutex` for thread safety
- Consider using `decimal.Decimal` for precise calculations
- Use `time.Time` for timestamps
- Implement proper validation
- Handle concurrent transfers safely
