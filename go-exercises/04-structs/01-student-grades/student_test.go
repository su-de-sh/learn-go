package main

import (
    "testing"
)

func TestNewStudent(t *testing.T) {
    student := NewStudent("John Doe")
    if student == nil {
        t.Error("NewStudent returned nil")
    }
    if student.Name != "John Doe" {
        t.Errorf("Expected name 'John Doe', got '%s'", student.Name)
    }
}

func TestAddGrade(t *testing.T) {
    student := NewStudent("John Doe")
    student.AddGrade(85.5)
    if student.GradeCount() != 1 {
        t.Errorf("Expected 1 grade, got %d", student.GradeCount())
    }
}

func TestAverage(t *testing.T) {
    tests := []struct {
        name     string
        grades   []float64
        expected float64
    }{
        {
            name:     "single grade",
            grades:   []float64{85.5},
            expected: 85.5,
        },
        {
            name:     "multiple grades",
            grades:   []float64{85.5, 92.0, 78.5},
            expected: 85.33,
        },
        {
            name:     "no grades",
            grades:   []float64{},
            expected: 0.0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            student := NewStudent("Test Student")
            for _, grade := range tt.grades {
                student.AddGrade(grade)
            }
            result := student.Average()
            if result != tt.expected {
                t.Errorf("Average() = %v, want %v", result, tt.expected)
            }
        })
    }
}

func TestHighestGrade(t *testing.T) {
    tests := []struct {
        name     string
        grades   []float64
        expected float64
    }{
        {
            name:     "multiple grades",
            grades:   []float64{85.5, 92.0, 78.5},
            expected: 92.0,
        },
        {
            name:     "no grades",
            grades:   []float64{},
            expected: 0.0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            student := NewStudent("Test Student")
            for _, grade := range tt.grades {
                student.AddGrade(grade)
            }
            result := student.HighestGrade()
            if result != tt.expected {
                t.Errorf("HighestGrade() = %v, want %v", result, tt.expected)
            }
        })
    }
}

func TestLowestGrade(t *testing.T) {
    tests := []struct {
        name     string
        grades   []float64
        expected float64
    }{
        {
            name:     "multiple grades",
            grades:   []float64{85.5, 92.0, 78.5},
            expected: 78.5,
        },
        {
            name:     "no grades",
            grades:   []float64{},
            expected: 0.0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            student := NewStudent("Test Student")
            for _, grade := range tt.grades {
                student.AddGrade(grade)
            }
            result := student.LowestGrade()
            if result != tt.expected {
                t.Errorf("LowestGrade() = %v, want %v", result, tt.expected)
            }
        })
    }
}

func TestGradeCount(t *testing.T) {
    student := NewStudent("Test Student")
    if student.GradeCount() != 0 {
        t.Errorf("Expected 0 grades, got %d", student.GradeCount())
    }

    student.AddGrade(85.5)
    if student.GradeCount() != 1 {
        t.Errorf("Expected 1 grade, got %d", student.GradeCount())
    }

    student.AddGrade(92.0)
    if student.GradeCount() != 2 {
        t.Errorf("Expected 2 grades, got %d", student.GradeCount())
    }
} 