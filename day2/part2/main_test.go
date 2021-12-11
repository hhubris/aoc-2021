package main

import (
	"testing"
)

func Test_position_move(t *testing.T) {

	type move struct {
		moveType string
		distance int
	}

	moves := []move{
		{
			moveType: "forward",
			distance: 5,
		},
		{
			moveType: "down",
			distance: 5,
		},
		{
			moveType: "forward",
			distance: 8,
		},
		{
			moveType: "up",
			distance: 3,
		},
		{
			moveType: "down",
			distance: 8,
		},
		{
			moveType: "forward",
			distance: 2,
		},
	}

	wants := []position{
		{
			h: 5,
			d: 0,
			a: 0,
		},
		{
			h: 5,
			d: 0,
			a: 5,
		},
		{
			h: 13,
			d: 40,
			a: 5,
		},
		{
			h: 13,
			d: 40,
			a: 2,
		},
		{
			h: 13,
			d: 40,
			a: 10,
		},
		{
			h: 15,
			d: 60,
			a: 10,
		},
	}

	p := position{}

	for i, m := range moves {
		p = p.move(m.moveType, m.distance)
		want := wants[i]

		if p != want {
			t.Errorf("after move %d, got %#v, want %#v", i, p, want)
		}
	}

}
