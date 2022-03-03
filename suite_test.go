package g

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLibs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Libs Suite")
}
