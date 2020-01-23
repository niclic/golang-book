package chapter6

func FindSmallest(nums []int) (s int) {
	s = nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < s {
			s = nums[i]
		}
	}
	return s
}
