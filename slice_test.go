package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Slice", func() {
	It("Copy should work", func() {
		Expect(Copy([]int{1, 2, 3})).To(Equal([]int{1, 2, 3}))
		Expect(Copy([]string{"a", "b", "c"})).To(Equal([]string{"a", "b", "c"}))
	})
})
