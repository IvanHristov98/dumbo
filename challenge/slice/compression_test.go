package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/IvanHristov98/dumbo/challenge/slice"
)

var _ = Describe("Compression", func() {
	DescribeTable("Compress", func(uncompressed, compressed string) {
		Expect(slice.Compress(uncompressed)).To(Equal(compressed))
	},
		Entry("empty string", "", ""),
		Entry("single char", "AAA", "A3"),
		Entry("complex compression", "AAABBCCCC", "A3B2C4"),
		Entry("senseless compression", "AABBCC", "AABBCC"),
	)
})
