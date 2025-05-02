package main

import "testing"

func TestBallsAreSortable(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected bool
	}{
		{
			name: "Balls are sortable",
			input: [][]int{
				{1, 2},
				{2, 1},
			},
			expected: true,
		},
		{
			name: "Balls are unsortable",
			input: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: false,
		},
		{
			name: "Balls are sortable",
			input: [][]int{
				{1},
			},
			expected: true,
		},
		{
			name: "Balls are sortable",
			input: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: true,
		},
		{
			name: "Balls are sortable",
			input: [][]int{
				{0, 1, 0},
				{1, 5, 2},
				{0, 2, 0},
			},
			expected: true,
		},
		{
			name: "Balls are unsortable",
			input: [][]int{
				{0, 0, 0, 0},
				{0, 5, 1, 0},
				{0, 0, 5, 0},
				{0, 0, 0, 0},
			},
			expected: true,
		},
		{
			name: "Balls are unsortable",
			input: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: false,
		},
		{
			name: "Balls are unsortable",
			input: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: false,
		},
		{
			name: "Balls are sortable",
			input: [][]int{
				{6, 0, 0},
				{0, 4, 0},
				{0, 0, 8},
			},
			expected: true,
		},
		{
			name: "Balls are sortable",
			input: [][]int{
				{3, 1, 2},
				{2, 4, 1},
				{1, 2, 8},
			},
			expected: true,
		},
		{
			name: "Balls are sortable",
			input: [][]int{
				{3, 5, 4, 19},
				{3, 4, 1, 13},
				{6, 5, 8, 0},
				{19, 5, 8, 5},
			},
			expected: true,
		},
		{
			name: "Balls are unsortable",
			input: [][]int{
				{3, 5, 1, 22},
				{3, 4, 1, 13},
				{6, 5, 8, 0},
				{22, 5, 8, 5},
			},
			expected: false,
		},
		{
			name: "Balls are unsortable",
			input: [][]int{
				{3, 5, 1, 19},
				{3, 4, 12, 13},
				{6, 5, 8, 0},
				{19, 5, 8, 5},
			},
			expected: false,
		},
		{
			name: "Balls are unsortable",
			input: [][]int{
				{3, 5, 4, 19, 2},
				{3, 4, 1, 13, 2},
				{6, 5, 8, 0, 8},
				{19, 5, 8, 5, 1},
				{3, 5, 1, 5, 1},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ballsAreSortable(test.input); got != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}
