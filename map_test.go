package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Map", func() {
	It("MergeMap should work", func() {
		m1 := map[string]string{
			"a": "b",
		}
		m2 := map[string]string{
			"c": "d",
		}

		Expect(MergeMap(m1, m2)).To(Equal(map[string]string{
			"a": "b",
			"c": "d",
		}))
	})
})
