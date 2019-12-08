package dp

type NumArray struct {
	nums    []int
	sumNums []int
}

func Constructor(nums []int) NumArray {

	newNums := make([]int, len(nums), cap(nums))
	copy(newNums, nums)
	sumNums := make([]int, len(newNums))
	if len(nums) == 0 {
		return NumArray{
			nums:    newNums,
			sumNums: sumNums,
		}
	}

	sumNums[0] = newNums[0]

	for i := 1; i < len(newNums); i++ {
		sumNums[i] = sumNums[i-1] + newNums[i]
	}

	return NumArray{
		nums:    newNums,
		sumNums: sumNums,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < j || i >= len(this.nums) || j >= len(this.nums) {
		return 0
	}
	return this.sumNums[j] - this.sumNums[i] + this.nums[i]
}
