package _59_Employee_Free_Time

import "testing"

func TestEmployeeFreeTime(t *testing.T) {
	t.Run("Default case 1", func(t *testing.T) {
		employees := make([][]*Interval, 3)
		employees[0] = []*Interval{{1, 2}, {5, 6}}
		employees[1] = []*Interval{{1, 3}}
		employees[2] = []*Interval{{4, 10}}
		actual := employeeFreeTime(employees)
		expected := []*Interval{{3, 4}}
		if len(actual) != len(expected) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("Default case 2", func(t *testing.T) {
		employees := make([][]*Interval, 3)
		employees[0] = []*Interval{{1, 3}, {6, 7}}
		employees[1] = []*Interval{{2, 4}}
		employees[2] = []*Interval{{2, 5}, {9, 12}}
		actual := employeeFreeTime(employees)
		expected := []*Interval{{5, 6}, {7, 8}}
		if len(actual) != len(expected) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
