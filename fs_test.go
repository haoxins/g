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

	It("CatFile should work", func() {
		s := CatFile("LICENSE")
		Expect(s).To(ContainSubstring("MIT License"))
		Expect(s).To(ContainSubstring("Copyright (c) 2021 Hao Xin"))
	})
})
