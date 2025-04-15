package main

import (
    "errors"
    "fmt"
    "math"
)

// Calculator represents a calculator with basic arithmetic operations
type Calculator struct {
    history []string
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
    return &Calculator{
        history: make([]string, 0),
    }
}

// Add adds two numbers
func (c *Calculator) Add(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, errors.New("invalid input: NaN values")
    }
    if math.IsInf(a, 0) || math.IsInf(b, 0) {
        return 0, errors.New("invalid input: infinite values")
    }
    result := a + b
    if math.IsInf(result, 0) {
        return 0, errors.New("result overflow")
    }
    c.history = append(c.history, fmt.Sprintf("%f + %f = %f", a, b, result))
    return result, nil
}

// Subtract subtracts two numbers
func (c *Calculator) Subtract(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, errors.New("invalid input: NaN values")
    }
    if math.IsInf(a, 0) || math.IsInf(b, 0) {
        return 0, errors.New("invalid input: infinite values")
    }
    result := a - b
    if math.IsInf(result, 0) {
        return 0, errors.New("result overflow")
    }
    c.history = append(c.history, fmt.Sprintf("%f - %f = %f", a, b, result))
    return result, nil
}

// Multiply multiplies two numbers
func (c *Calculator) Multiply(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, errors.New("invalid input: NaN values")
    }
    if math.IsInf(a, 0) || math.IsInf(b, 0) {
        return 0, errors.New("invalid input: infinite values")
    }
    result := a * b
    if math.IsInf(result, 0) {
        return 0, errors.New("result overflow")
    }
    c.history = append(c.history, fmt.Sprintf("%f * %f = %f", a, b, result))
    return result, nil
}

// Divide divides two numbers
func (c *Calculator) Divide(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, errors.New("invalid input: NaN values")
    }
    if math.IsInf(a, 0) || math.IsInf(b, 0) {
        return 0, errors.New("invalid input: infinite values")
    }
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    result := a / b
    if math.IsInf(result, 0) {
        return 0, errors.New("result overflow")
    }
    c.history = append(c.history, fmt.Sprintf("%f / %f = %f", a, b, result))
    return result, nil
}

// Power calculates a raised to the power of b
func (c *Calculator) Power(base, exponent float64) (float64, error) {
    if math.IsNaN(base) || math.IsNaN(exponent) {
        return 0, errors.New("invalid input: NaN values")
    }
    if math.IsInf(base, 0) || math.IsInf(exponent, 0) {
        return 0, errors.New("invalid input: infinite values")
    }
    if base < 0 && !math.IsInf(exponent, 0) && exponent != float64(int64(exponent)) {
        return 0, errors.New("invalid operation: negative base with non-integer exponent")
    }
    result := math.Pow(base, exponent)
    if math.IsInf(result, 0) {
        return 0, errors.New("result overflow")
    }
    c.history = append(c.history, fmt.Sprintf("%f ^ %f = %f", base, exponent, result))
    return result, nil
}

// GetHistory returns the operation history
func (c *Calculator) GetHistory() []string {
    return c.history
} 