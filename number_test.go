package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Number", func() {
	It("ForceInt should work", func() {
		var s = "123"
		Expect(ForceInt(s, 0)).To(Equal(123))
		s = "abc"
		Expect(ForceInt(s, -1)).To(Equal(-1))
	})

	It("ForceInt32 should work", func() {
		var s = "123"
		Expect(ForceInt32(s, 0)).To(Equal(int32(123)))
		s = "abc"
		Expect(ForceInt32(s, -1)).To(Equal(int32(-1)))
	})

	It("ForceInt64 should work", func() {
		var s = "123"
		Expect(ForceInt64(s, 0)).To(Equal(int64(123)))
		s = "abc"
		Expect(ForceInt64(s, -1)).To(Equal(int64(-1)))
	})
})
