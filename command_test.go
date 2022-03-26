package g

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Command", func() {
	It("Run should work", func() {
		out, err := Run("ls", "-a")
		Expect(err).To(BeNil())
		Expect(strings.Contains(out, "command.go")).To(BeTrue())
		Expect(strings.Contains(out, "command_test.go")).To(BeTrue())

		out, err = Run("xxoo")
		Expect(err).NotTo(BeNil())
		Expect(out).NotTo(Equal(""))
	})
})
