# Go Programming Concepts Demo

This project demonstrates several fundamental concepts in Go programming language through a simple shape area calculator.

## Concepts Demonstrated

### 1. Interfaces

- Go interfaces define a set of methods that a type must implement
- In this code, the `Shape` interface requires two methods:
  ```go
  type Shape interface {
      Area() (float32, error)
      Introduce() string
  }
  ```

### 2. Error Handling

- Go uses explicit error handling through the `error` interface
- Methods can return both a value and an error
- The code demonstrates a generic error handling function:
  ```go
  func handleError[T Shape](value T, er error) T {
      if er != nil {
          log.Fatalf("Error: %v", er)
      }
      return value
  }
  ```
- This generic function provides a clean way to handle errors for any type that implements the `Shape` interface
- Example usage:
  ```go
  square := handleError(square.NewSquare(3))
  ```

### 3. Generics

- Go 1.18+ supports generic programming
- The `handleError` function demonstrates generic type parameters:
  ```go
  func handleError[T Shape](value T, er error) T
  ```
- This allows the function to work with any type that implements the `Shape` interface

### 4. Structs and Methods

- Go uses structs to group related data
- Methods can be attached to structs using receiver syntax
- Example:
  ```go
  type Circle struct {
      radius float32
  }

  func (c Circle) Area() (float32, error) {
      if c.radius < 0 {
          return 0, errors.New("radius must be positive")
      }
      return PI * c.radius * c.radius, nil
  }
  ```

### 5. Type Implementation

- Types implicitly implement interfaces by implementing all required methods
- No explicit declaration is needed
- `Circle`, `Rectangle`, and `Square` all implement the `Shape` interface

### 6. Constants

- Constants are declared using the `const` keyword
- Example:
  ```go
  const PI = 3.14
  ```

### 7. String Formatting

- Using `fmt.Printf` for formatted output
- Format specifiers:
  - `%.2f` for float with 2 decimal places
  - `%-40s` for left-aligned string with 40-character width

### 8. Slices and Loops

- Using slices to store multiple values
- `for` loop with range to iterate over slices
- Example:
  ```go
  shapes := []Shape{circle, rectangle, square}
  for _, shape := range shapes {
      // ...
  }
  ```

### 9. Package Management

- Go uses a module-based package management system
- The project is organized into multiple packages:
  - `main`: Contains the entry point and core logic
  - `circle`: Implements circle shape functionality
  - `rectangle`: Implements rectangle shape functionality
  - `square`: Implements square shape functionality
- Dependencies are managed through `go.mod`:
  ```go
  module github.com/pubestpubest/go101area

  go 1.21
  ```
- Each shape package is self-contained and follows Go's package naming conventions
- Importing packages:
  ```go
  import (
      "github.com/pubestpubest/go101area/circle"
      "github.com/pubestpubest/go101area/rectangle"
      "github.com/pubestpubest/go101area/square"
  )
  ```

## Running the Code

```bash
go run main.go
```

## Output Example

```
Circle with radius: 1.00                            area: 3.14
Rectangle with width: 2.00, height: 2.00            area: 4.00
Square with width: 3.00                             area: 9.00
```

## Key Takeaways

1. Go's interface system is powerful and implicit
2. Generics provide type-safe code reuse
3. Methods can be attached to any type
4. String formatting in Go is flexible and supports alignment
5. Go's type system is simple yet effective
6. Slices and loops make it easy to work with collections of data
7. Generic error handling can make code more concise and maintainable
