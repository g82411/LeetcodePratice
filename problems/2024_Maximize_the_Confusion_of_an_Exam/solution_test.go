package _024_Maximize_the_Confusion_of_an_Exam

import "testing"

func TestMaxConsecutiveAnswers(t *testing.T) {
	testcases := []struct {
		answerKey string
		k         int
		except    int
	}{
		{
			"TTFF",
			2,
			4,
		},
		{
			"TFFT",
			1,
			3,
		},
		{
			"TTFTTFTT",
			1,
			5,
		},
	}
	for _, testcase := range testcases {
		actual := maxConsecutiveAnswers(testcase.answerKey, testcase.k)
		if actual != testcase.except {
			t.Errorf("answerKey: %v, k: %v, except: %v, actual: %v", testcase.answerKey, testcase.k, testcase.except, actual)
		}
	}
}
