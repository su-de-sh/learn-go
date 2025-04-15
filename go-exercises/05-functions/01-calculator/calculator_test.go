package main

import (
    "math"
    "testing"
)

func TestNewCalculator(t *testing.T) {
    calc := NewCalculator()
    if calc == nil {
        t.Error("NewCalculator returned nil")
    }
    if len(calc.history) != 0 {
        t.Error("New calculator should have empty history")
    }
}

func TestAdd(t *testing.T) {
    calc := NewCalculator()
    tests := []struct {
        name     string
        a, b     float64
        expected float64
        wantErr  bool
    }{
        {"basic addition", 2, 3, 5, false},
        {"negative numbers", -2, -3, -5, false},
        {"zero", 0, 0, 0, false},
        {"large numbers", 1e308, 1e308, 0, true}, // overflow
        {"NaN", math.NaN(), 1, 0, true},
        {"Inf", math.Inf(1), 1, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := calc.Add(tt.a, tt.b)
            if tt.wantErr {
                if err == nil {
                    t.Error("expected error but got none")
                }
                return
            }
            if err != nil {
                t.Errorf("unexpected error: %v", err)
                return
            }
            if result != tt.expected {
                t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestSubtract(t *testing.T) {
    calc := NewCalculator()
    tests := []struct {
        name     string
        a, b     float64
        expected float64
        wantErr  bool
    }{
        {"basic subtraction", 5, 3, 2, false},
        {"negative result", 3, 5, -2, false},
        {"zero", 0, 0, 0, false},
        {"large numbers", -1e308, 1e308, 0, true}, // overflow
        {"NaN", math.NaN(), 1, 0, true},
        {"Inf", math.Inf(1), 1, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := calc.Subtract(tt.a, tt.b)
            if tt.wantErr {
                if err == nil {
                    t.Error("expected error but got none")
                }
                return
            }
            if err != nil {
                t.Errorf("unexpected error: %v", err)
                return
            }
            if result != tt.expected {
                t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestMultiply(t *testing.T) {
    calc := NewCalculator()
    tests := []struct {
        name     string
        a, b     float64
        expected float64
        wantErr  bool
    }{
        {"basic multiplication", 2, 3, 6, false},
        {"negative numbers", -2, -3, 6, false},
        {"zero", 0, 5, 0, false},
        {"large numbers", 1e154, 1e154, 0, true}, // overflow
        {"NaN", math.NaN(), 1, 0, true},
        {"Inf", math.Inf(1), 1, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := calc.Multiply(tt.a, tt.b)
            if tt.wantErr {
                if err == nil {
                    t.Error("expected error but got none")
                }
                return
            }
            if err != nil {
                t.Errorf("unexpected error: %v", err)
                return
            }
            if result != tt.expected {
                t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestDivide(t *testing.T) {
    calc := NewCalculator()
    tests := []struct {
        name     string
        a, b     float64
        expected float64
        wantErr  bool
    }{
        {"basic division", 6, 2, 3, false},
        {"fraction result", 5, 2, 2.5, false},
        {"zero numerator", 0, 5, 0, false},
        {"zero denominator", 5, 0, 0, true},
        {"NaN", math.NaN(), 1, 0, true},
        {"Inf", math.Inf(1), 1, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := calc.Divide(tt.a, tt.b)
            if tt.wantErr {
                if err == nil {
                    t.Error("expected error but got none")
                }
                return
            }
            if err != nil {
                t.Errorf("unexpected error: %v", err)
                return
            }
            if result != tt.expected {
                t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestPower(t *testing.T) {
    calc := NewCalculator()
    tests := []struct {
        name     string
        base, exp float64
        expected float64
        wantErr  bool
    }{
        {"basic power", 2, 3, 8, false},
        {"negative exponent", 2, -1, 0.5, false},
        {"zero exponent", 5, 0, 1, false},
        {"negative base with non-integer exponent", -2, 1.5, 0, true},
        {"NaN", math.NaN(), 1, 0, true},
        {"Inf", math.Inf(1), 1, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := calc.Power(tt.base, tt.exp)
            if tt.wantErr {
                if err == nil {
                    t.Error("expected error but got none")
                }
                return
            }
            if err != nil {
                t.Errorf("unexpected error: %v", err)
                return
            }
            if result != tt.expected {
                t.Errorf("Power(%v, %v) = %v, want %v", tt.base, tt.exp, result, tt.expected)
            }
        })
    }
}

func TestHistory(t *testing.T) {
    calc := NewCalculator()
    calc.Add(2, 3)
    calc.Multiply(4, 5)
    
    history := calc.GetHistory()
    if len(history) != 2 {
        t.Errorf("expected history length 2, got %d", len(history))
    }
    if history[0] != "2.000000 + 3.000000 = 5.000000" {
        t.Errorf("unexpected history entry: %s", history[0])
    }
    if history[1] != "4.000000 * 5.000000 = 20.000000" {
        t.Errorf("unexpected history entry: %s", history[1])
    }
} 