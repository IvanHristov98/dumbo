package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/IvanHristov98/dumbo/challenge/slice"
)

var _ = Describe("Rotation", func() {
	DescribeTable("IsRotation", func(a, b string, expected bool) {
		Expect(slice.IsRotation(a, b)).To(Equal(expected))
	},
		Entry("empty strings", "", "", true),
		Entry("rotated strings", "rotated", "atedrot", true),
		Entry("not rotated strings", "banana", "nanaab", false),
	)
})
