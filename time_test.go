package g

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Time", func() {
	It("TimeConverter should work", func() {
		ts := "2022-03-04T16:56:27+08:00"
		var tc = &TimeConverter{}
		t := tc.FromString(ts, time.RFC3339)
		Expect(tc.ToString(t, time.RFC3339)).To(Equal(ts))
	})

	It("GetTime should work", func() {
		ts := "2022-03-04T16:56:27+08:00"
		t := GetTime(ts, time.RFC3339)
		Expect(t.Unix()).To(Equal(int64(1646384187)))
	})

	It("GetCurrentTime should work", func() {
		Expect(GetTime(GetCurrentTime(time.RFC3339), time.RFC3339).Unix() > 1646384187).To(BeTrue())
	})

	It("SetCurrentTime should work", func() {
		target := ""
		SetCurrentTime(&target, time.RFC3339)
		Expect(GetTime(target, time.RFC3339).Unix() > 1646384187).To(BeTrue())
	})
})
