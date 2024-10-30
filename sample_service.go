// Sample-service is a simple program with lots of functions to test GoDI against
package main

import (
	"time"

	rookout "github.com/Rookout/GoSDK"
)

func main() {
	rookout.Start(rookout.RookOptions{
		Token:  "f5729394eba7028b8a0ccf55069e25ba5fb33014a6524cc20c8dc0a117517481",
		Labels: map[string]string{"env": "dev"},
	})

	ticker := time.NewTicker(time.Second / 2)
	for range ticker.C {
		ExecuteOther()
		ExecuteBasicFuncs()
		ExecuteMultiParamFuncs()
		ExecuteStringFuncs()
		ExecuteArrayFuncs()
		ExecuteSliceFuncs()
		ExecuteStructFuncs()
		ExecuteStackAndInlining()
		ExecutePointerFuncs()

		// unsupported for MVP, should not cause crashes
		ExecuteGenericFuncs()
		ExecuteMapFuncs()
		ExecuteInterfaceFuncs()
		go Return_goroutine_id()
	}
}
