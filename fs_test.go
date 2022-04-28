package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test fs", func() {
	It("Copy should work", func() {
		err := CopyFile("fs.go", "fs.go.out")
		Expect(err).To(BeNil())
	})
})
