package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndexFromCoordinate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		coordinate string
		want       int
		want2      int
	}{
		// TODO: Add test cases.
		struct {
			name       string
			coordinate string
			want       int
			want2      int
		}{
			name:       "Success",
			coordinate: "a4",
			want:       3,
			want2:      0,
		},
		struct {
			name       string
			coordinate string
			want       int
			want2      int
		}{
			name:       "Success",
			coordinate: "d2",
			want:       1,
			want2:      3,
		},
		{
			name:       "Success",
			coordinate: "a2",
			want:       1,
			want2:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row, col := GetIndexFromCoordinate(tt.coordinate)

			assert.Equal(t, tt.want, row)
			assert.Equal(t, tt.want2, col)
		})
	}
}
