package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test URL", func() {
	It("EncodeURL should work", func() {
		encoded := EncodeURL(`https://www.google.com?components=["search"]`)
		Expect(encoded).To(Equal("https://www.google.com?components=%5B%22search%22%5D"))
	})
})
