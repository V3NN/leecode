package __twoSum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	r := twoSum(nums, target)
	fmt.Println(r)

	r = twoSumByHash(nums, target)
	fmt.Println(r)

}
