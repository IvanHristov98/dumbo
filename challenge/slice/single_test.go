package slice_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/IvanHristov98/dumbo/challenge/slice"
)

var _ = Describe("Single", func() {
	Context("when provided input is correct", func() {
		DescribeTable("FindDiffChar", func(a, b string, extectedDiff rune) {
			diff, err := slice.FindDiffChar(a, b)

			Expect(diff).To(Equal(extectedDiff))
			Expect(err).To(BeNil())
		},
			Entry("when first is shorter and 1 rune", "a", "aa", 'a'),
			Entry("when first is shorter with multiple runes", "abcd", "abcde", 'e'),
			Entry("when second is longer and 1 rune", "aa", "a", 'a'),
			Entry("when second is longer with multiple runes", "abcde", "abcd", 'e'),
			Entry("when the characters are shuffled and repeating", "abbcccdde", "babcdccd", 'e'),
		)
	})

	Context("when provided input is incorrect", func() {
		Context("and there are no different runes", func() {
			It("returns an error", func() {
				diff, err := slice.FindDiffChar("abcd", "abcd")

				Expect(diff).To(Equal(rune(0)))
				Expect(err).To(Equal(fmt.Errorf("No different runes")))
			})
		})

		Context("and there are more than one different runes", func() {
			It("returns an error", func() {
				diff, err := slice.FindDiffChar("abcdef", "abcd")

				Expect(diff).To(Equal(rune(0)))
				Expect(err).To(Equal(fmt.Errorf("More than one different runes")))
			})
		})
	})
})
