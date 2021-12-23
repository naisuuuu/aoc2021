package main

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOverlaps(t *testing.T) {
	tests := []struct {
		a, b Cube
		want bool
	}{
		{
			a:    Cube{0, 4, 0, 4, 0, 4},
			b:    Cube{3, 4, 3, 4, 3, 4},
			want: true,
		},
		{
			a:    Cube{0, 4, 0, 4, 0, 4},
			b:    Cube{5, 6, 3, 4, 3, 4},
			want: false,
		},
		{
			a:    Cube{0, 4, 0, 4, 0, 4},
			b:    Cube{4, 5, 3, 4, 3, 4},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := overlaps(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		a, b Cube
		want bool
	}{
		{
			a:    Cube{0, 4, 0, 4, 0, 4},
			b:    Cube{3, 4, 3, 4, 3, 4},
			want: true,
		},
		{
			a:    Cube{3, 4, 3, 4, 3, 4},
			b:    Cube{0, 4, 0, 4, 0, 4},
			want: false,
		},
		{
			a:    Cube{0, 4, 0, 4, 0, 4},
			b:    Cube{4, 5, 3, 4, 3, 4},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := contains(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		a, b Cube
		want []Cube
	}{
		{
			a: Cube{0, 4, 0, 4, 0, 4},
			b: Cube{2, 2, 2, 2, 2, 2},
			want: []Cube{
				{0, 1, 0, 4, 0, 4},
				{3, 4, 0, 4, 0, 4},
				{2, 2, 0, 1, 0, 4},
				{2, 2, 3, 4, 0, 4},
				{2, 2, 2, 2, 0, 1},
				{2, 2, 2, 2, 3, 4},
			},
		},
		{
			a:    Cube{0, 4, 0, 4, 0, 4},
			b:    Cube{0, 4, 0, 4, 0, 4},
			want: nil,
		},
		{
			a: Cube{0, 4, 0, 4, 0, 4},
			b: Cube{3, 6, 0, 4, 0, 4},
			want: []Cube{
				{0, 2, 0, 4, 0, 4},
			},
		},
	}
	for i, tt := range tests {
		got := split(tt.a, tt.b)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
