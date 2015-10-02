package assert

import (
	"fmt"
	"runtime"
)

// Assert panics if the passed-in condition is not met.
func Assert(cond bool) {
	if cond {
		return
	}

	_, file, line, ok := runtime.Caller(1)
	if ok {
		panic(fmt.Sprintf("ASSERT: Condition not met at %s:%d", file, line))
	}

	panic("ASSERT: Condition not met.")
}
