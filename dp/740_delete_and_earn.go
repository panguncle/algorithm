package dp

/*
DeleteAndEarn 740
if take nums[i] =>
1. can take all of its copies.
2. can't take any of copies of nums[i-1] and nums[i + 1]

=> reduced to 198 Hourse Robber
=> Hourses[i] has all the copies of num whose value is i
[2, 2. 3, 3, 3, 4] => [0, 2 * 2, 3 * 3, 4] =>
Time: O(n + r) reduction + O(r) solving rob = O(n+r)
Space: O(r)
r = max{nums}
*/
func DeleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	maxV := maxN(nums...)

	scores := make([]int, maxV+1)
	for _, num := range nums {
		scores[num]++
	}

	prev1 := scores[0]
	prev2 := maxN(scores[0], scores[1])
	for i := 2; i < len(scores); i++ {
		cur := maxN(prev1+scores[i]*i, prev2)
		prev1 = prev2
		prev2 = cur
	}

	return prev2
}

func maxN(items ...int) int {

	var maxV = items[0]
	for i := 1; i < len(items); i++ {
		if items[i] > maxV {
			maxV = items[i]
		}
	}

	return maxV
}
