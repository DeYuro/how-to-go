package main

import "golang.org/x/exp/constraints"

func main() {
	s, ss := "a", []string{"a", "b"}
	i, is := 1, []int{1, 2}

	println(contains(s, ss))
	println(contains(i, is))

	a, b := 1, 20
	println(max(a, b))
	as, bs := "lorem", "dorem"
	println(max(as, bs))

	printErr(InvalidNameError)
	printErr(ResourceNotFoundError)
	printErr(EmptyValueError)
	printErr(ServiceError)
}

func contains[T comparable](value T, sliceOfValue []T) bool {
	for _, v := range sliceOfValue {
		if v == value {
			return true
		}
	}

	return false
}

type CanCompare interface {
	constraints.Ordered
}

func max[T CanCompare](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Err string
type InternalErr string

const (
	InvalidNameError Err = `InvalidName`
	EmptyValueError  Err = `EmptyValue`
)

const (
	ResourceNotFoundError InternalErr = `ResourceNotFound`
	ServiceError          InternalErr = `ServiceError`
)

type AllErr interface {
	Err | InternalErr
}

func printErr[T AllErr](error T) {
	println(error)
}
