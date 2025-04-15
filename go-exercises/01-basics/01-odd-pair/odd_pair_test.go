package main

import (
    "reflect"
    "testing"
)

func TestOddPair(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected []int
    }{
        {
            name:     "returns first two odd numbers",
            input:    []int{1, 2, 3, 4, 5},
            expected: []int{1, 3},
        },
        {
            name:     "returns single odd number when only one exists",
            input:    []int{10, 15, 20},
            expected: []int{15},
        },
        {
            name:     "returns empty slice when no odd numbers",
            input:    []int{2, 4, 6, 8},
            expected: []int{},
        },
        {
            name:     "handles empty input slice",
            input:    []int{},
            expected: []int{},
        },
        {
            name:     "handles nil input slice",
            input:    nil,
            expected: []int{},
        },
        {
            name:     "handles slice with all odd numbers",
            input:    []int{1, 3, 5, 7, 9},
            expected: []int{1, 3},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := OddPair(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("OddPair(%v) = %v, want %v", tt.input, result, tt.expected)
            }
        })
    }
} 