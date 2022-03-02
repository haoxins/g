package golibs

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMap(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Map Suite")
}

func TestMergeMap(t *testing.T) {
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
}
