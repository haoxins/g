package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test URL", func() {
	It("EncodeURL should work", func() {
		encoded, err := EncodeURL(`https://www.google.com?components=["search"]`)
		Expect(err).To(BeNil())
		Expect(encoded).To(Equal("https://www.google.com?components=%5B%22search%22%5D"))
	})

	It("EncodeURL should work with empty string", func() {
		encoded, err := EncodeURL("")
		Expect(err).To(BeNil())
		Expect(encoded).To(Equal(""))
	})

	It("GetURLHost should work", func() {
		host, err := GetURLHost(`https://www.google.com?components=["search"]`)
		Expect(err).To(BeNil())
		Expect(host).To(Equal("www.google.com"))

		host, err = GetURLHost("")
		Expect(err).To(BeNil())
		Expect(host).To(Equal(""))
	})
})
