package golibs

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test String", func() {
	It("IsBlank should work", func() {
		var s = "		"
		Expect(IsBlank(&s)).To(BeTrue())
		s = "	"
		Expect(IsBlank(&s)).To(BeTrue())
		s = "  "
		Expect(IsBlank(&s)).To(BeTrue())
		s = " "
		Expect(IsBlank(&s)).To(BeTrue())
		s = ""
		Expect(IsBlank(&s)).To(BeTrue())
		Expect(IsBlank(nil)).To(BeTrue())
		s = "no"
		Expect(IsBlank(&s)).To(BeFalse())
	})
})
