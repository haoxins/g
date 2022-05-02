package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Map", func() {
	It("MergeMap with [string]string should work", func() {
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

	It("MergeMap with [int]int should work", func() {
		m1 := map[int]int{
			1: 1,
		}
		m2 := map[int]int{
			2: 2,
		}

		Expect(MergeMap(m1, m2)).To(Equal(map[int]int{
			1: 1,
			2: 2,
		}))
	})
})
