package dividenconquer

/*
MajorityElement : 169
@see 摩尔投票法: 前提, 必须存在超过半数的值
https://www.cnblogs.com/grandyang/p/4233501.html
*/
func MajorityElement(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	maxK := nums[0]
	maxCount := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != maxK {
			maxCount--
		} else {
			maxCount++
		}

		if maxCount == 0 {
			maxK = nums[i]
			maxCount = 1
		}
	}

	return maxK
}
