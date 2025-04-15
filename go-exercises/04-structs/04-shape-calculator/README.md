# Shape Calculator

Implement a shape calculator system using structs and interfaces.

## Requirements

1. Create an interface `Shape` with methods:

   - `Area() float64`: Calculate area
   - `Perimeter() float64`: Calculate perimeter
   - `String() string`: String representation

2. Implement the following structs that satisfy the `Shape` interface:

   - `Rectangle`:

     - `Width`: Rectangle width
     - `Height`: Rectangle height

   - `Circle`:

     - `Radius`: Circle radius

   - `Triangle`:

     - `Base`: Triangle base
     - `Height`: Triangle height
     - `Sides`: Lengths of all three sides

   - `Square`:
     - `Side`: Square side length

3. Create a `ShapeCalculator` struct with methods:
   - `CalculateArea(shape Shape) float64`: Calculate area of any shape
   - `CalculatePerimeter(shape Shape) float64`: Calculate perimeter of any shape
   - `CompareShapes(shape1, shape2 Shape) string`: Compare two shapes by area
   - `ScaleShape(shape Shape, factor float64) Shape`: Scale a shape by factor
   - `IsValid(shape Shape) bool`: Validate shape dimensions

## Example Usage

```go
// Create shapes
rect := &Rectangle{Width: 5, Height: 3}
circle := &Circle{Radius: 2.5}
triangle := &Triangle{Base: 4, Height: 3, Sides: []float64{3, 4, 5}}
square := &Square{Side: 4}

// Create calculator
calc := &ShapeCalculator{}

// Calculate areas
rectArea := calc.CalculateArea(rect)
circleArea := calc.CalculateArea(circle)
fmt.Println(rectArea)   // 15.0
fmt.Println(circleArea) // 19.634954084936208

// Calculate perimeters
rectPerimeter := calc.CalculatePerimeter(rect)
circlePerimeter := calc.CalculatePerimeter(circle)
fmt.Println(rectPerimeter)   // 16.0
fmt.Println(circlePerimeter) // 15.707963267948966

// Compare shapes
comparison := calc.CompareShapes(rect, circle)
fmt.Println(comparison) // "Circle has larger area than Rectangle"

// Scale shape
scaledRect := calc.ScaleShape(rect, 2.0)
fmt.Println(scaledRect) // "Rectangle{Width: 10, Height: 6}"

// Validate shape
isValid := calc.IsValid(triangle)
fmt.Println(isValid) // true
```

## Requirements

1. The implementation should handle:

   - Invalid dimensions
   - Zero or negative values
   - Floating-point precision
   - Shape-specific calculations
   - Interface compliance

2. Implementation details:
   - Use proper error handling
   - Handle edge cases
   - Consider numerical precision
   - Implement efficient algorithms
   - Use appropriate math functions

## Tips

- Use `math.Pi` for circle calculations
- Use `math.Sqrt` for square root
- Consider using `float64` for precision
- Implement proper validation
- Handle degenerate cases
- Use appropriate rounding
