package g

import (
	"fmt"
)

func Assert(a, b any) {
	if a != b {
		panic(fmt.Errorf("Assert %v to equal %v", a, b))
	}
}
