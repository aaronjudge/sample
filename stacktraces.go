package main

import "fmt"

//nolint:all
//go:noinline
func stack_A() {
	stack_B()
}

//nolint:all
//go:noinline
func stack_B() {
	stack_C()
}

//nolint:all
//go:noinline
func stack_C() string {
	return fmt.Sprintf("hello %d!", 1)
}

//nolint:all
//go:noinline
func call_inlined_func_chain() {
	inline_me_1()
}

//nolint:all
func inline_me_1() {
	inline_me_2()
}

//nolint:all
func inline_me_2() {
	inline_me_3()
}

//nolint:all
func inline_me_3() {
	not_inlined()
}

//nolint:all
//go:noinline
func not_inlined() string {
	return fmt.Sprintf("hello %d!", 42)
}

//nolint:all
func ExecuteStackAndInlining() {
	stack_A()
	call_inlined_func_chain()
}
