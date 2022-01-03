# gogm

The `gogm` package is a graphics math library using the new language feature of Go: generics. The API is currently unstable under development. Therefore, do not use this API for production yet. The API will stabilize when Go 1.18 is released.

## Installation

This module uses Go 1.18, which is still in beta. The beta can be installed using these commands:

```sh
go install golang.org/dl/go1.18beta1@latest

go1.18beta1 download
```

Go 1.18 can now be used by replacing `go` with `go1.18beta1`. Example:

```sh
go1.18beta1 run .
```

Optionally, you can update `gopls` to work with Go 1.18:

```sh
go1.18beta1 install golang.org/x/tools/gopls@latest
```

Finally, you can add `gogm` to your module:

```sh
go1.18beta1 get github.com/oyberntzen/gogm
```

## Usage

Vectors are initialized with a type argument, specifying the underlying type of the vector. The type can be all floats and integers.

```go
// Vector with float32 as underlying type
pos := gogm.Vec3[float32]{1.0, 2.0, 3.0}

// Vector with uint8 as underlying type
red := gogm.Vec3[uint8]{255, 0, 0}
```

No new vectors are declared by `gogm`. Therefore, the results of all operations are stored in the struct that called the method.

```go
v1 := gogm.Vec3[int]{}
v2 := gogm.Vec3[int]{1, 2, 3}
v3 := gogm.Vec3[int]{4, 5, 6}

// The result will be stored in v1
v1.Add(&v2, &v3)

// The resulting vector can be used in the operation
v2.Add(&v2, &v3)

// Exeptions are when the operation returns a single value
result := v2.Cross(&v3)
```

## Documentation

Go 1.18 packages does not work on [pkg.go.dev](https://pkg.go.dev/) yet, but this will likely be fixed when Go 1.18 is officially released. For now, you have to stick with reading the source files.
