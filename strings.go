package main

//nolint:all
//go:noinline
func test_single_string(x string) {}

//nolint:all
func ExecuteStringFuncs() {
	test_single_string("abc")
}
