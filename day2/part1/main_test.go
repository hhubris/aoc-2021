package main

import (
	"testing"
)

func Test_position_move(t *testing.T) {
	tests := []struct {
		name     string
		moveType string
		distance int
		want     position
	}{
		{
			name:     "horizontal",
			moveType: "forward",
			distance: 7,
			want:     position{h: 7, d: 0},
		},
		{
			name:     "down",
			moveType: "down",
			distance: 3,
			want:     position{h: 0, d: 3},
		},
		{
			name:     "up",
			moveType: "up",
			distance: 9,
			want:     position{h: 0, d: -9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := position{}
			got := p.move(tt.moveType, tt.distance)
			if tt.want != got {
				t.Errorf("want: %#v, got %#v", tt.want, got)
			}
		})
	}
}
