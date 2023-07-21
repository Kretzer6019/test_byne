package clients

import "testing"

func TestGetValue(t *testing.T) {
	isEven := true
	isOdd := false

	if isEven {
		value := GetValue(isEven)
		if value%2 != 0 {
			t.Error("O valor deveria ser par")
		}
	} else if isOdd {
		value := GetValue(isOdd)
		if value%2 == 0 {
			t.Error("O valor deveria ser ímpar")
		}
	}
}
