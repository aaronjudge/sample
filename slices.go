package main

import "fmt"

//nolint:all
//go:noinline
func doNothing(x []int) {}

//nolint:all
//go:noinline
func sprintSlice(x []int) string {
	return fmt.Sprintf("%v\n", x)
}

//nolint:all
//go:noinline
func expandSlice(x []int) {
	x = append(x, []int{9, 10, 11, 12}...)
	doNothing(x)
}

//nolint:all
//go:noinline
func test_uint_slice(u []uint) {}

//nolint:all
//go:noinline
func test_struct_slice(xs []structWithNoStrings) {}

//nolint:all
func ExecuteSliceFuncs() {
	originalSlice := []int{1, 2, 3}
	expandSlice(originalSlice)
	sprintSlice(originalSlice)

	test_uint_slice([]uint{1, 2, 3})
	test_struct_slice([]structWithNoStrings{{42, true}, {24, true}})
}
