package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/IvanHristov98/dumbo/challenge/slice"
)

var _ = Describe("Reverse", func() {
	DescribeTable("Reverse", func(straight, reversed string) {
		Expect(slice.Reverse(straight)).To(Equal(reversed))
	},
		Entry("empty string", "", ""),
		Entry("reverse", "foo bar", "rab oof"),
	)
})
