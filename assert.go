package g

import (
	"fmt"
)

func Assert(ok bool, msg string) {
	if !ok {
		panic(fmt.Errorf(msg))
	}
}
