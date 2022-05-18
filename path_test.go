package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Path", func() {
	It("Resolve should work", func() {
		Expect(Resolve("/a/b", "../c")).To(Equal("/a/c"))
		Expect(Resolve("/a/b", "./c")).To(Equal("/a/b/c"))
		Expect(Resolve("/a/b", "c")).To(Equal("/a/b/c"))
		Expect(Resolve("/a/b", "/c")).To(Equal("/c"))
		Expect(Resolve("", "/c")).To(Equal("/c"))
		Expect(Resolve("", "c")).To(HaveSuffix("/g/c"))
		Expect(Resolve("/", "c")).To(Equal("/c"))
	})

	It("Resolve URL should work", func() {
		url, err := ResolveURL("https://example.com", "/a/b")
		Expect(err).To(BeNil())
		Expect(url).To(Equal("https://example.com/a/b"))

		url, err = ResolveURL("https://example.com/a", "/b")
		Expect(err).To(BeNil())
		Expect(url).To(Equal("https://example.com/b"))

		url, err = ResolveURL("https://example.com/a/", "b")
		Expect(err).To(BeNil())
		Expect(url).To(Equal("https://example.com/a/b"))

		url, err = ResolveURL("https://example.com/a", "b")
		Expect(err).To(BeNil())
		Expect(url).To(Equal("https://example.com/a/b"))

		url, err = ResolveURL("", "c")
		Expect(err).To(BeNil())
		Expect(url).To(Equal("/c"))
	})
})
