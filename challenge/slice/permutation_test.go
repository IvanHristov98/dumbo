package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/IvanHristov98/dumbo/challenge/slice"
)

var _ = Describe("Permutation", func() {
	DescribeTable("IsPermutation", func(a, b string, expected bool) {
		Expect(slice.IsPermutation(a, b)).To(Equal(expected))
	},
		Entry("empty strings", "", "", true),
		Entry("permutation", "cat", "act", true),
		Entry("no permutation with less characters", "cat", "ac", false),
		Entry("no permutation with different characters", "cat", "actd", false),
		Entry("case sensitivity", "Cat", "act", false),
	)
})
