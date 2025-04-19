# TestGen

TestGen is a command-line utility written in Go that automatically generates test files with simple body for a provided Go file.

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
