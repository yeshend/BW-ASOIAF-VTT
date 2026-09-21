package dice

import "testing"

func TestIsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		shade    Shade
		expected bool
	}{
		// Black Shade (4+)
		{name: "Black 1", value: 1, shade: Black, expected: false},
		{name: "Black 2", value: 2, shade: Black, expected: false},
		{name: "Black 3", value: 3, shade: Black, expected: false},
		{name: "Black 4", value: 4, shade: Black, expected: true},
		{name: "Black 5", value: 5, shade: Black, expected: true},
		{name: "Black 6", value: 6, shade: Black, expected: true},

		// Gray Shade (3+)
		{name: "Gray 1", value: 1, shade: Gray, expected: false},
		{name: "Gray 2", value: 2, shade: Gray, expected: false},
		{name: "Gray 3", value: 3, shade: Gray, expected: true},
		{name: "Gray 4", value: 4, shade: Gray, expected: true},
		{name: "Gray 5", value: 5, shade: Gray, expected: true},
		{name: "Gray 6", value: 6, shade: Gray, expected: true},

		// White Shade (2+)
		{name: "White 1", value: 1, shade: White, expected: false},
		{name: "White 2", value: 2, shade: White, expected: true},
		{name: "White 3", value: 3, shade: White, expected: true},
		{name: "White 4", value: 4, shade: White, expected: true},
		{name: "White 5", value: 5, shade: White, expected: true},
		{name: "White 6", value: 6, shade: White, expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isSuccess(test.value, test.shade)
			if actual != test.expected {
				t.Errorf("Expected: %v, got: %v", test.expected, actual)
			}
		})
	}
}
