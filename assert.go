package g

import (
	"errors"
)

func Assert(ok bool, msg string) {
	if !ok {
		panic(errors.New(msg))
	}
}
