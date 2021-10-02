package slice_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/IvanHristov98/dumbo/challenge/slice"
)

var _ = Describe("Sum", func() {
	Describe("TwoSum", func() {
		Context("when two indices sum up to a provided value", func() {
			It("returns the indices", func() {
				nums := []int{1, 3, 2, -7, 5}
				val := 7

				expectedIndices := []uint{2, 4}
				indices, err := slice.TwoSum(nums, val)

				Expect(indices).To(Equal(expectedIndices))
				Expect(err).To(BeNil())
			})
		})

		DescribeTable("when two indices don't sum up to a provided value", func(nums []int, val int) {
			expectedIndices := []uint{}
			expectedErr := fmt.Errorf("No matching indices found")
			indices, err := slice.TwoSum(nums, val)

			Expect(indices).To(Equal(expectedIndices))
			Expect(err).To(Equal(expectedErr))
		},
			Entry("and a number is equal of half of the searched value", []int{1, 2}, 4),
			Entry("and the array is with less than two elements", []int{}, 4),
			Entry("and the array has no matching indices", []int{1, 2, 3, 4}, 10),
		)
	})
})
