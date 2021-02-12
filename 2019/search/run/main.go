package main

import (
	"fmt"

	"github.com/panguncle/algorithm/2019/search"
)

func main() {
	testGenerateParentheses()
}

func testCombine() {
	ret := search.Combination(4, 2)
	fmt.Println(ret)
}

func testCombineSum() {
	ret := search.CombinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(ret)
}

func testCombineSum2() {
	ret := search.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(ret)
}

func testLetterCombination() {
	ret := search.LetterCombinations("23")
	fmt.Println(ret)
}

func testPermutation() {
	ret := search.Permutation([]int{1, 2, 3})
	fmt.Println(ret)
}

func testPermutationII() {
	ret := search.PermuteUnique([]int{1, 1, 2})
	fmt.Println(ret)
}

func testSubsets() {
	ret := search.Subsets([]int{1, 2})
	fmt.Println(ret)
}

func testCombinationSumIII() {
	ret := search.CombinationSumIII(3, 7)
	fmt.Println(ret)
	ret = search.CombinationSumIII(3, 9)
	fmt.Println(ret)
}

func testLetterCasePermutation() {
	ret := search.LetterCasePermutation("a1b2")
	fmt.Println(ret)
}

func testGenerateParentheses() {
	ret := search.GenerateParentheses(3)
	fmt.Println(ret)
}
