# Student Grades

Create a `Student` struct and implement methods to manage student grades.

## Requirements

1. Create a `Student` struct with the following fields:

   - `Name` (string)
   - `Grades` ([]float64)

2. Implement the following methods:

   - `AddGrade(grade float64)`: Adds a grade to the student's grades
   - `Average() float64`: Returns the average grade
   - `HighestGrade() float64`: Returns the highest grade
   - `LowestGrade() float64`: Returns the lowest grade
   - `GradeCount() int`: Returns the number of grades

3. Create a constructor function `NewStudent(name string) *Student`

## Example Usage

```go
student := NewStudent("John Doe")
student.AddGrade(85.5)
student.AddGrade(92.0)
student.AddGrade(78.5)

fmt.Println(student.Average())     // 85.0
fmt.Println(student.HighestGrade()) // 92.0
fmt.Println(student.LowestGrade())  // 78.5
fmt.Println(student.GradeCount())   // 3
```

## Requirements

1. The `Student` struct should be properly encapsulated
2. All methods should handle edge cases (no grades, invalid grades)
3. The `Average()` method should return 0.0 if there are no grades
4. The `HighestGrade()` and `LowestGrade()` methods should handle empty grades
5. Grades should be stored with 2 decimal places precision

## Tips

- Use pointer receivers for methods that modify the student's state
- Consider using `math.Max()` and `math.Min()` for grade comparisons
- Use `float64` for precise grade calculations
- Remember to handle edge cases in all methods
