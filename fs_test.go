package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test fs", func() {
	It("Exists should work", func() {
		exists, err := Exists("fs.go")
		Expect(err).To(BeNil())
		Expect(exists).To(BeTrue())
		exists, err = Exists("fs_.go")
		Expect(err).To(BeNil())
		Expect(exists).To(BeFalse())
	})

	It("CopyFile should work", func() {
		err := CopyFile("fs.go", "fs.go.out")
		Expect(err).To(BeNil())
	})

	It("SyncFile should work", func() {
		err := SyncFile("fs.go", "fs/fs.go.out", true)
		Expect(err).To(BeNil())
	})

	It("ReadFile should work", func() {
		s := ReadFile("LICENSE")
		Expect(s).To(ContainSubstring("MIT License"))
		Expect(s).To(ContainSubstring("Copyright (c) 2021 Hao Xin"))
	})
})
