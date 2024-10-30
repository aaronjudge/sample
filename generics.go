package main

type typeWithGenerics[V comparable] struct {
	Value V
}

//nolint:all
//go:noinline
func (x typeWithGenerics[V]) Guess(value V) bool {
	return x.Value == value
}

//nolint:all
func ExecuteGenericFuncs() {
	x := typeWithGenerics[string]{Value: "generics work"}
	x.Guess("generics work")

	y := typeWithGenerics[int]{Value: 42}
	y.Guess(21)
}
