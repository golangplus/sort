# sort
Plus to standard sort package.

[Godoc](http://godoc.org/github.com/golangplus/sort)

## Featured
```go
// InterfaceStruct is a struct implementing sort.Interface given closures
type InterfaceStruct struct {...}

// SortF calls sort.Sort by closures. Since Interface.Len always returns a constant,
// it is an int parameter rather than a closure here.
func SortF(Len int, Less func(i, j int) bool, Swap func(i, j int)) {...}
```

## LICENSE
BSD license
