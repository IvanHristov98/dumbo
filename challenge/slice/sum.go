package slice

import "fmt"

func TwoSum(nums []int, val int) ([]uint, error) {
	indicesByVal := map[int]uint{}

	for idx, num := range nums {
		subtractedVal := val - num

		if otherIdx, exists := indicesByVal[subtractedVal]; exists {
			return []uint{otherIdx, uint(idx)}, nil
		}

		setDefaultValIdx(num, uint(idx), indicesByVal)
	}

	return []uint{}, fmt.Errorf("No matching indices found")
}

func setDefaultValIdx(num int, idx uint, indicesByVal map[int]uint) {
	if _, exists := indicesByVal[num]; !exists {
		indicesByVal[num] = idx
	}
}
