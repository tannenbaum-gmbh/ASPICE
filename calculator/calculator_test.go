package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 5, 9},
		{"negative numbers", -3, -5, -8},
		{"mixed numbers", -3, 5, 2},
		{"zeros", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 8, 3, 5},
		{"negative numbers", -8, -3, -5},
		{"mixed numbers", 8, -3, 11},
		{"zeros", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Subtract(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Subtract(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 5, 15},
		{"negative numbers", -3, -5, 15},
		{"mixed numbers", -3, 5, -15},
		{"zeros", 0, 5, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Multiply(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Multiply(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expected    int
		expectError bool
	}{
		{"positive numbers", 15, 3, 5, false},
		{"negative numbers", -15, -3, 5, false},
		{"mixed numbers", -15, 3, -5, false},
		{"division by zero", 15, 0, 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Divide(tc.a, tc.b)

			if tc.expectError {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tc.a, tc.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tc.a, tc.b, err)
				}
				if result != tc.expected {
					t.Errorf("Divide(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
				}
			}
		})
	}
}
