package g

import (
	"fmt"
)

func Assert(a, b any) {
	if a != b {
		panic(fmt.Errorf("Assert failed! %v [not equals] %v", a, b))
	}
}
