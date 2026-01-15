package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test string", func() {
	It("LowerTrim should work", func() {
		Expect(LowerTrim(" Hi ")).To(Equal("hi"))
	})
})
