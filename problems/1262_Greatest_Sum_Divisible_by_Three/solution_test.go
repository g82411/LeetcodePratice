package _262_Greatest_Sum_Divisible_by_Three

import "testing"

func TestMaxSumDivThree(t *testing.T) {
	testcase := []struct {
		nums   []int
		except int
	}{
		{
			nums:   []int{3, 6, 5, 1, 8},
			except: 18,
		},
		{
			nums:   []int{4},
			except: 0,
		},
		{
			nums:   []int{1, 2, 3, 4, 4},
			except: 12,
		},
	}
	for _, tc := range testcase {
		actual := maxSumDivThree(tc.nums)
		if actual != tc.except {
			t.Errorf("nums: %v, except: %v, actual: %v", tc.nums, tc.except, actual)
		}
	}

}
