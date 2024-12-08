package main

import (
	"testing"
)

func TestNextMul(t *testing.T) {
	
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5)) mul ( 8 , 5 ) mul(1234,2) mul(123,1233)"
	tests := []struct {
		expectedLeft    int
		expectedRight int
	}{
		{ expectedLeft: 2, expectedRight: 4},
		{ expectedLeft: 5, expectedRight: 5},
		{ expectedLeft: 11, expectedRight: 8},
		{ expectedLeft: 8, expectedRight: 5},
	}
	tokenizer := NewTokenizer(input)
	t.Log(tokenizer)

	for _, tt := range tests {
		mul := tokenizer.nextMul()
		
		t.Log("left:", mul.LeftNumber, " - right:", mul.RightNumber)
		if mul.LeftNumber != tt.expectedLeft {
			t.Fatal("expected:", tt.expectedLeft, " - actual:", mul.LeftNumber)
		}

		if mul.RightNumber != tt.expectedRight {
			t.Fatal("expected:", tt.expectedRight, " - actual:", mul.RightNumber)
		}
	}
}
