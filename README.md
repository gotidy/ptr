# ptr [![GoDoc](https://godoc.org/github.com/gotidy/ptr?status.svg)](https://godoc.org/github.com/gotidy/ptr)

`ptr` contains functions for simplified creation of pointers from constants of basic types.

## Installation

`go get github.com/gotidy/ptr`

## Examples

This code:

```go
p := ptr.Int(10)
```

is the equivalent for:

```go
i := int(10)
p := &i  
```

## Documentation

[GoDoc](http://godoc.org/github.com/gotidy/ptr)

## License

[Apache 2.0](https://github.com/gotidy/ptr/blob/master/LICENSE)