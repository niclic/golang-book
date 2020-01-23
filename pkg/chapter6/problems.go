package chapter6

func FindSmallest(nums []int) (s int) {
	s = nums[0]
	for _, i := range nums {
		if i < s {
			s = i
		}
	}
	return s
}
