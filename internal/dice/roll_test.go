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

func TestCountSuccesses(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		shade    Shade
		expected int
	}{
		// Black Shade (4+)
		{name: "Black empty", values: []int{}, shade: Black, expected: 0},
		{name: "Black 0 in 3", values: []int{1, 2, 3}, shade: Black, expected: 0},
		{name: "Black 3 in 3", values: []int{4, 5, 6}, shade: Black, expected: 3},
		{name: "Black 2 in 4", values: []int{1, 4, 2, 6}, shade: Black, expected: 2},

		// Gray Shade (3+)
		{name: "Gray empty", values: []int{}, shade: Gray, expected: 0},
		{name: "Gray 1 in 3", values: []int{1, 2, 3}, shade: Gray, expected: 1},
		{name: "Gray 3 in 3", values: []int{4, 5, 6}, shade: Gray, expected: 3},
		{name: "Gray 2 in 4", values: []int{1, 4, 2, 6}, shade: Gray, expected: 2},

		// White Shade (2+)
		{name: "White empty", values: []int{}, shade: White, expected: 0},
		{name: "White 2 in 3", values: []int{1, 2, 3}, shade: White, expected: 2},
		{name: "White 3 in 3", values: []int{4, 5, 6}, shade: White, expected: 3},
		{name: "White 3 in 4", values: []int{1, 4, 2, 6}, shade: White, expected: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := countSuccesses(test.values, test.shade)
			if actual != test.expected {
				t.Errorf("Expected: %v, got: %v", test.expected, actual)
			}
		})
	}
}
