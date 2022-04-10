package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Path", func() {
	It("Resolve should work", func() {
		Expect(Resolve("", "c")).To(HaveSuffix("/g/c"))
		Expect(Resolve("/a/b", "./c")).To(Equal("/a/b/c"))
		Expect(Resolve("/a/b", "../c")).To(Equal("/a/c"))
		Expect(Resolve("/a/b", "c")).To(Equal("/a/b/c"))
		Expect(Resolve("/a/b", "/c")).To(Equal("/c"))
		Expect(Resolve("", "/c")).To(Equal("/c"))
	})
})
