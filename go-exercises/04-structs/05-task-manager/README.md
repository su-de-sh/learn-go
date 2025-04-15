# Task Manager

Implement a task management system using structs and methods.

## Requirements

1. Create the following structs:

   - `Task`:

     - `ID`: Unique task identifier
     - `Title`: Task title
     - `Description`: Task description
     - `Status`: Task status (pending/in-progress/completed)
     - `Priority`: Task priority (low/medium/high)
     - `DueDate`: Task due date
     - `CreatedAt`: Task creation time
     - `UpdatedAt`: Last update time
     - `AssignedTo`: Assignee information
     - `Tags`: Slice of task tags

   - `User`:

     - `ID`: Unique user identifier
     - `Name`: User name
     - `Email`: User email
     - `Tasks`: Slice of assigned task IDs
     - `Role`: User role (admin/user)

   - `TaskManager`:
     - `Tasks`: Map of tasks by ID
     - `Users`: Map of users by ID
     - `Tags`: Map of tasks by tag

2. Implement the following methods:
   - `NewTaskManager() *TaskManager`: Create new task manager
   - `AddTask(task *Task) error`: Add new task
   - `AddUser(user *User) error`: Add new user
   - `AssignTask(taskID, userID string) error`: Assign task to user
   - `UpdateTaskStatus(taskID, status string) error`: Update task status
   - `GetTask(taskID string) (*Task, error)`: Get task by ID
   - `GetUserTasks(userID string) ([]*Task, error)`: Get user's tasks
   - `GetTasksByTag(tag string) []*Task`: Get tasks by tag
   - `GetOverdueTasks() []*Task`: Get overdue tasks
   - `SearchTasks(query string) []*Task`: Search tasks by title/description
   - `FilterTasks(status, priority string) []*Task`: Filter tasks by status/priority

## Example Usage

```go
// Create task manager
manager := NewTaskManager()

// Add user
user := &User{
    ID:    "U001",
    Name:  "John Doe",
    Email: "john@example.com",
    Role:  "user",
}
manager.AddUser(user)

// Add task
task := &Task{
    ID:          "T001",
    Title:       "Implement login feature",
    Description: "Create user authentication system",
    Status:      "pending",
    Priority:    "high",
    DueDate:     time.Now().AddDate(0, 0, 7),
    CreatedAt:   time.Now(),
    UpdatedAt:   time.Now(),
    Tags:        []string{"auth", "frontend"},
}
manager.AddTask(task)

// Assign task
manager.AssignTask("T001", "U001")

// Update status
manager.UpdateTaskStatus("T001", "in-progress")

// Get user's tasks
tasks, _ := manager.GetUserTasks("U001")
fmt.Println(tasks) // [Task{ID: "T001", Title: "Implement login feature", ...}]

// Get tasks by tag
authTasks := manager.GetTasksByTag("auth")
fmt.Println(authTasks) // [Task{ID: "T001", Title: "Implement login feature", ...}]

// Search tasks
results := manager.SearchTasks("login")
fmt.Println(results) // [Task{ID: "T001", Title: "Implement login feature", ...}]

// Filter tasks
highPriority := manager.FilterTasks("", "high")
fmt.Println(highPriority) // [Task{ID: "T001", Title: "Implement login feature", ...}]
```

## Requirements

1. The implementation should handle:

   - Task status transitions
   - User assignments
   - Due dates
   - Task filtering
   - Search functionality
   - Error conditions

2. Implementation details:
   - Use proper error handling
   - Implement thread safety
   - Handle edge cases
   - Consider data consistency
   - Implement efficient search
   - Handle time zones

## Tips

- Use `sync.Mutex` for thread safety
- Use maps for O(1) lookups
- Implement proper validation
- Handle concurrent access
- Consider using indexes for search
- Use `time.Time` for dates
- Consider using enums for status/priority
