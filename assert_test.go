package g

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Test Assert", func() {
	It("Should work", func() {
		Assert(123 == 123, "Should equal")
	})

	It("Should work", func() {
		println("TODO - panic")
	})
})
