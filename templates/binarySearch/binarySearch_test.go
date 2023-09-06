package binarySearch

import "testing"

func TestFindTarget(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 3
		expect := 2
		actual := findTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("when target not in nums return -1", func(t *testing.T) {
		nums := []int{1, 2, 4, 5}
		target := 3
		expect := -1
		actual := findTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("only target", func(t *testing.T) {
		nums := []int{1}
		target := 1
		expect := 0
		actual := findTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("only one element w/o target", func(t *testing.T) {
		nums := []int{0}
		target := 1
		expect := -1
		actual := findTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}

func TestFindFirstTarget(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 3
		expect := 2
		actual := findFirstTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("when target not in nums return -1", func(t *testing.T) {
		nums := []int{1, 2, 4, 5}
		target := 3
		expect := -1
		actual := findFirstTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("only target", func(t *testing.T) {
		nums := []int{1}
		target := 1
		expect := 0
		actual := findFirstTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("only one element w/o target", func(t *testing.T) {
		nums := []int{0}
		target := 1
		expect := -1
		actual := findFirstTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("when duplicate, find first element", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
		target := 1
		expect := 0
		actual := findFirstTarget(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}

func TestListLowerBound(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 3
		expect := 2
		actual := listLowerBound(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("when target not in nums return first element not less than target", func(t *testing.T) {
		nums := []int{1, 2, 4, 5}
		target := 3
		expect := 2
		actual := listLowerBound(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("when target large than all of then target", func(t *testing.T) {
		nums := []int{1, 2, 4, 5}
		target := 6
		expect := 4
		actual := listLowerBound(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}

func TestListUpperBound(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		target := 3
		expect := 3
		actual := listUpperBound(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("when target not in nums return first element large than target", func(t *testing.T) {
		nums := []int{1, 2, 4, 5}
		target := 3
		expect := 2
		actual := listUpperBound(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("when target large than all of then target", func(t *testing.T) {
		nums := []int{1, 2, 4, 5}
		target := 6
		expect := 4
		actual := listUpperBound(nums, target)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
