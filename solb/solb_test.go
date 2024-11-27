package solb

import "testing"

// TestDecodeToMinSumNumber ทดสอบฟังก์ชัน decodeToMinSumNumber
func TestDecodeToMinSumNumber(t *testing.T) {
	tests := []struct {
		encoded  string
		expected string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
		{"L", "10"}, // Case: simple L
		{"R", "01"}, // Case: simple R
		{"=", "00"}, // Case: simple =
		// {"L=R", "101"},
		// {"RRLL=", "012210"},
	}

	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			result := decodeToMinSumNumber(test.encoded)
			if result != test.expected {
				t.Errorf("For input '%s', expected '%s' but got '%s'", test.encoded, test.expected, result)
			}
		})
	}
}
