# Go Programming Concepts Demo

This project demonstrates several fundamental concepts in Go programming language through a simple shape area calculator.

## Concepts Demonstrated

### 1. Interfaces

- Go interfaces define a set of methods that a type must implement
- In this code, the `Area` interface requires two methods:
  ```go
  type Area interface {
      Area() (float32, error)
      Introduce() string
  }
  ```

### 2. Error Handling

- Go uses explicit error handling through the `error` interface
- Methods can return both a value and an error
- Example of error handling in the code:
  ```go
  func (c Circle) Area() (float32, error) {
      if c.radius < 0 {
          return 0, errors.New("radius must be positive")
      }
      return PI * c.radius * c.radius, nil
  }
  ```
- The main function demonstrates error handling:
  ```go
  area, error := shape.Area()
  if error != nil {
      fmt.Println("Error: ", error)
      return
  }
  ```

### 2. Structs and Methods

- Go uses structs to group related data
- Methods can be attached to structs using receiver syntax
- Example:

  ```go
  type Circle struct {
      radius float32
  }

  func (c Circle) Area() float32 {
      return PI * c.radius * c.radius
  }
  ```

### 3. Type Implementation

- Types implicitly implement interfaces by implementing all required methods
- No explicit declaration is needed
- `Circle`, `Rectangle`, and `Square` all implement the `Area` interface

### 4. Constants

- Constants are declared using the `const` keyword
- Example:
  ```go
  const PI = 3.14
  ```

### 5. String Formatting

- Using `fmt.Sprintf` for string formatting
- Format specifiers:
  - `%.2f` for float with 2 decimal places
  - `%-40s` for left-aligned string with 40-character width

### 6. Slices and Loops

- Using slices to store multiple values
- `for` loop with range to iterate over slices
- Example:
  ```go
  shapes := []Area{circle, rec, squre}
  for _, shape := range shapes {
      // ...
  }
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
2. Methods can be attached to any type
3. String formatting in Go is flexible and supports alignment
4. Go's type system is simple yet effective
5. Slices and loops make it easy to work with collections of data
