# go-util

`go-util` is a lightweight collection of Go utilities for slice and map operations, struct reflection, and caller-name inspection.

## Requirements

- Go 1.14 or later

## Installation

```bash
go get github.com/lokixg7/go-util
```

Import the packages you need:

```go
import (
	"github.com/lokixg7/go-util/array"
	"github.com/lokixg7/go-util/maps"
	"github.com/lokixg7/go-util/runtime"
	"github.com/lokixg7/go-util/structs"
)
```

## Slice Utilities

Package: `github.com/lokixg7/go-util/array`

### `InArray`

Reports whether a target value exists in a slice.

```go
found := array.InArray(int64(7), []int64{1, 3, 5, 7})
// found == true
```

### `Intersect`

Calculates the intersection of two slices and writes it to the supplied result-slice pointer.

```go
var result []int64
err := array.Intersect(
	[]int64{1, 2, 3, 4},
	[]int64{1, 3, 5},
	&result,
)
// result == []int64{1, 3}
```

### `Diff`

Calculates the difference of `X` relative to `Y`, retaining elements that occur only in `X`.

```go
var result []int64
err := array.Diff(
	[]int64{1, 2, 3, 4},
	[]int64{2, 3, 5, 6},
	&result,
)
// result == []int64{1, 4}
```

### `Unique`

Removes duplicate slice elements while preserving their first-occurrence order.

```go
var result []int64
err := array.Unique([]int64{1, 2, 2, 3, 1}, &result)
// result == []int64{1, 2, 3}
```

### `Explode`

Joins slice elements with a delimiter.

```go
joined := array.Explode(",", []string{"go", "util", "array"})
// joined == "go,util,array"
```

> Elements passed to `InArray`, `Intersect`, `Diff`, and `Unique` must be comparable, such as strings, numbers, booleans, and pointers. Non-comparable values such as slices, maps, and functions are not supported. `Intersect`, `Diff`, and `Unique` use JSON encoding and decoding to write results, so the result parameter must be a pointer to a JSON-decodable slice.

## Map Utilities

Package: `github.com/lokixg7/go-util/maps`

### `Map2List`

Extracts the values of a map into a slice.

```go
var result []string
err := maps.Map2List(map[string]string{"a": "go", "b": "util"}, &result)
// result contains "go" and "util"
```

> `Map2List` uses JSON encoding and decoding to write results, so the result parameter must be a pointer to a JSON-decodable slice.

## Struct Utilities

Package: `github.com/lokixg7/go-util/structs`

### `SetStructFields`

Sets struct fields in bulk by field name. The first argument must be a struct pointer. Supported field types are `string`, `int64`, `uint64`, `float64`, and `bool`.

```go
type User struct {
	Name string
	Age  int64
}

user := &User{Name: "Andrew", Age: 26}
err := structs.SetStructFields(user, map[string]interface{}{
	"Name": "Blue",
	"Age":  int64(27),
})
// user == &User{Name: "Blue", Age: 27}
```

### `ConvertToMap`

Converts a slice of structs or struct pointers to a map using the specified field as the key.

```go
type User struct {
	ID   uint64
	Name string
}

users := []User{{ID: 1, Name: "Blue"}, {ID: 2, Name: "Crank"}}
byID, err := structs.ConvertToMap(users, "ID")
// byID[uint64(1)] == User{ID: 1, Name: "Blue"}
```

## Call Stack Utilities

Package: `github.com/lokixg7/go-util/runtime`

- `GetCurCalleeFunc()`: returns the name of its direct caller.
- `GetParentCallFunc()`: returns the name of its caller's parent.

```go
func example() string {
	return runtime.GetCurCalleeFunc()
}

// example() == "example"
```

## Testing

Run all tests from the repository root:

```bash
go test ./...
```

Run `TestInArray` in the array package:

```bash
go test -v ./array -run '^TestInArray$'
```

Run `TestMap2List` in the maps package:

```bash
go test -v ./maps -run '^TestMap2List$'
```
