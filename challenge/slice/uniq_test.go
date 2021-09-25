package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/IvanHristov98/dumbo/challenge/slice"
)

var _ = Describe("Uniq", func() {
	DescribeTable("HasUniqueChars", func(str string, expected bool) {
		Expect(slice.HasUniqueChars(str)).To(Equal(expected))
	},
		Entry("empty string", "", true),
		Entry("string with unique chars", "bar", true),
		Entry("string with repeating chars", "foo", false),
	)
})
