# Student Grades

Implement a student grade management system using maps.

## Requirements

1. Create a `GradeManager` struct with the following methods:

   - `AddStudent(name string) error`: Add a new student
   - `AddGrade(name, subject string, grade float64) error`: Add a grade
   - `GetGrades(name string) (map[string]float64, error)`: Get all grades for a student
   - `GetAverage(name string) (float64, error)`: Get student's average grade
   - `GetSubjectAverage(subject string) (float64, error)`: Get class average for a subject
   - `GetTopStudents(n int) []StudentRank`: Get top n students by average
   - `GetSubjectTopStudents(subject string, n int) []StudentRank`: Get top n students in a subject
   - `RemoveStudent(name string) error`: Remove a student

2. Each method should:
   - Handle invalid grades
   - Handle non-existent students
   - Handle non-existent subjects
   - Be efficient for large datasets

## Example Usage

```go
manager := NewGradeManager()

// Add students
manager.AddStudent("John Doe")
manager.AddStudent("Jane Smith")

// Add grades
manager.AddGrade("John Doe", "Math", 85.5)
manager.AddGrade("John Doe", "Science", 92.0)
manager.AddGrade("Jane Smith", "Math", 88.0)
manager.AddGrade("Jane Smith", "Science", 95.0)

// Get student grades
grades, err := manager.GetGrades("John Doe")
fmt.Println(grades) // map[Math:85.5 Science:92.0]

// Get student average
avg, err := manager.GetAverage("John Doe")
fmt.Println(avg) // 88.75

// Get subject average
subjectAvg, err := manager.GetSubjectAverage("Math")
fmt.Println(subjectAvg) // 86.75

// Get top students
top := manager.GetTopStudents(2)
fmt.Println(top) // [{Jane Smith 91.5} {John Doe 88.75}]

// Get top students in subject
topMath := manager.GetSubjectTopStudents("Math", 2)
fmt.Println(topMath) // [{Jane Smith 88.0} {John Doe 85.5}]
```

## Requirements

1. The implementation should handle:

   - Invalid grades (negative, >100)
   - Non-existent students
   - Non-existent subjects
   - Performance considerations
   - Memory usage

2. Implementation details:
   - Use efficient data structures
   - Validate grades
   - Handle edge cases
   - Consider space complexity

## Tips

- Use nested maps for student-subject-grade relationships
- Consider using a custom type for grades
- Use `sort` package for ranking students
- Consider using `big.Float` for precise calculations
- Handle concurrent access safely
