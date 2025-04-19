# TestGen

TestGen is a command-line utility written in Go that automatically generates test files with simple body for a provided Go file.

# Features

- Automatically generates test file (`*_test.go`) for a given Go source file.
- Creates test function templates for all __exported functions and methods__ found in the file.
- Determines the appropriate test structure based on the complexity of the function:
    - For simple logic: generates a basic test function.
    - For functions with multiple branches: generates a table-driven test function.
- When generating tests for methods of a struct, it analyzes the struct’s fields:
    - For any field that is of interface usertype, it automatically generates field initialization by mock.

# Installation

You can install `TestGen` using `go install`

```
go install github.com/jj-mon/testgen@latest
```

Make sure you have Go installed on your system and your `$GOPATH/bin` is in your PATH to use the `testgen` command globally.

# Usage

To generate tests for a Go file, use the following command:

```
testgen <path-to-your-go-file>
```
