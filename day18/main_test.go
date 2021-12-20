package main

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseTree(t *testing.T) {
	tests := []struct {
		in   string
		want *Node
	}{
		{
			in: "[1,2]",
			want: &Node{
				Left:  &Node{Value: 1},
				Right: &Node{Value: 2},
			},
		},
		{
			in: "[[1,2],3]",
			want: &Node{
				Left: &Node{
					Left:  &Node{Value: 1},
					Right: &Node{Value: 2},
				},
				Right: &Node{Value: 3},
			},
		},
		{
			in: "[9,[8,7]]",
			want: &Node{
				Left: &Node{Value: 9},
				Right: &Node{
					Left:  &Node{Value: 8},
					Right: &Node{Value: 7},
				},
			},
		},
		{
			in: "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := ParseTree(tt.in)
			if diff := cmp.Diff(tt.want, got); diff != "" && tt.want != nil {
				t.Errorf("ParseTree() mismatch (-want, +got):\n%s", diff)
			}
			gotStr := got.String()
			if gotStr != tt.in {
				t.Errorf("ParseTree() string: want %s, got %s", tt.in, gotStr)
			}
		})
	}
}

func TestExplode(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "[[[[[9,8],1],2],3],4]",
			want: "[[[[0,9],2],3],4]",
		},
		{
			in:   "[7,[6,[5,[4,[3,2]]]]]",
			want: "[7,[6,[5,[7,0]]]]",
		},
		{
			in:   "[[6,[5,[4,[3,2]]]],1]",
			want: "[[6,[5,[7,0]]],3]",
		},
		{
			in:   "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			want: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			in:   "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			want: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			tree := ParseTree(tt.in)
			tree.explode(4)
			if tt.want != tree.String() {
				t.Errorf("Explode() want %s, got %s", tt.want, tree)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			tree := ParseTree(tt.in)
			tree.reduce()
			if tt.want != tree.String() {
				t.Errorf("Reduce() want %s, got %s", tt.want, tree)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		in   []string
		want string
	}{
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
			},
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			},
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			},
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			in: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			in: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want: "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var ns []*Node
			for _, i := range tt.in {
				ns = append(ns, ParseTree(i))
			}
			got := Add(ns[0], ns[1:]...)
			if tt.want != got.String() {
				t.Errorf("Add() want %s, got %s", tt.want, got)
			}
		})
	}
}
