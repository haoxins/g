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

	It("Exists should work", func() {
		exists, err := Exists("fs.go")
		Expect(err).To(BeNil())
		Expect(exists).To(BeTrue())
		exists, err = Exists("fs_.go")
		Expect(err).To(BeNil())
		Expect(exists).To(BeFalse())
	})
})
