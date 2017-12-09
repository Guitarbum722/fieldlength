package qualified

import (
	"testing"
)

var splitCases = []struct {
	input, sep, qual string
	expected         int // num of resulting elements
}{
	{
		"one,two,three,four",
		",",
		"",
		4,
	},
	{
		"\"one,two\",three,four",
		",",
		"\"",
		3,
	},
	{
		"one",
		",",
		"",
		1,
	},
	{
		"one|two|three|four",
		"|",
		"",
		4,
	},
}

var fieldLenCases = []struct {
	input, sep, qual string
	expected         map[int]int
}{
	{
		"one,two,three,four",
		",",
		"",
		map[int]int{0: 3, 1: 3, 2: 5, 3: 4},
	},
	{
		"\"one,two\",three,four",
		",",
		"\"",
		map[int]int{0: 9, 1: 5, 2: 4},
	},
	{
		"one",
		",",
		"",
		map[int]int{0: 3},
	},
	{
		"one|two|three|four",
		"|",
		"",
		map[int]int{1: 3, 2: 5, 3: 4, 0: 3},
	},
}

func TestSplitWithQual(t *testing.T) {
	for _, tt := range splitCases {
		got := SplitWithQual(tt.input, tt.sep, tt.qual)

		if len(got) != tt.expected {
			t.Fatalf("SplitWithQual(%v) = len(%v) ; want %v", tt.input, len(got), tt.expected)
		}
	}
}

func TestFieldLengths(t *testing.T) {
	for _, tt := range fieldLenCases {
		got := FieldLengths(tt.input, tt.sep, tt.qual)
		for k := range tt.expected {
			if got[k] != tt.expected[k] {
				t.Fatalf("FieldLengths(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		}
	}
}
