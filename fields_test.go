package qualified

import "testing"

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

func TestSplitWithQual(t *testing.T) {
	for _, tt := range splitCases {
		got := SplitWithQual(tt.input, tt.sep, tt.qual)

		if len(got) != tt.expected {
			t.Fatalf("SplitWithQual(%v) = len(%v) ; want %v", tt.input, len(got), tt.expected)
		}
	}
}
