package TwoSum

import (
	"testing"
)

type paras struct {
	nums   []int
	target int
}

type answer struct {
	ans []int
}

type question struct {
	paras
	answer
}

func TestTwoSum(t *testing.T) {
	questions := []question{
		{
			paras{[]int{3, 2, 4}, 6},
			answer{[]int{1, 2}},
		},
		{
			paras{[]int{3, 2, 4}, 5},
			answer{[]int{0, 1}},
		},

		{
			paras{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			answer{[]int{1, 3}},
		},

		{
			paras{[]int{0, 1}, 1},
			answer{[]int{0, 1}},
		},

		{
			paras{[]int{0, 3}, 5},
			answer{[]int{}},
		},
	}

	for _, question := range questions {
		p, a := question.paras, question.answer

		// Equality Compare of TwoSum1
		result1 := TwoSum1(p.nums, p.target)
		if !compareArray(a.ans, result1) {
			t.Fail()
		}
		// Equality Compare of TwoSum2
		result2 := TwoSum2(p.nums, p.target)
		if !compareArray(a.ans, result2) {
			t.Fail()
		}
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	questions := []question{
		{
			paras{[]int{3, 2, 4}, 6},
			answer{[]int{1, 2}},
		},
		{
			paras{[]int{3, 2, 4}, 5},
			answer{[]int{0, 1}},
		},

		{
			paras{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			answer{[]int{1, 3}},
		},

		{
			paras{[]int{0, 1}, 1},
			answer{[]int{0, 1}},
		},

		{
			paras{[]int{0, 3}, 5},
			answer{[]int{}},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, question := range questions {
			p, _ := question.paras, question.answer
			TwoSum1(p.nums, p.target)
		}
	}
	b.StopTimer()
}

func BenchmarkTwoSum2(b *testing.B) {
	questions := []question{
		{
			paras{[]int{3, 2, 4}, 6},
			answer{[]int{1, 2}},
		},
		{
			paras{[]int{3, 2, 4}, 5},
			answer{[]int{0, 1}},
		},

		{
			paras{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			answer{[]int{1, 3}},
		},

		{
			paras{[]int{0, 1}, 1},
			answer{[]int{0, 1}},
		},

		{
			paras{[]int{0, 3}, 5},
			answer{[]int{}},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, question := range questions {
			p, _ := question.paras, question.answer
			TwoSum2(p.nums, p.target)
		}
	}
	b.StopTimer()
}

func compareArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}
