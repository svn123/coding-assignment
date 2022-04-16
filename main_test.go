package main

import (
	"fmt"
	"reflect"
	"testing"
)

//This code test the overall functionality of all the functions together
func TestFind_the_word(t *testing.T) {
	type args struct {
		puzzle [][]byte
		word   string
	}

	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				puzzle: [][]byte{
					{'a', 'b', 'c', 'd'},
					{'e', 'f', 'g', 'h'},
					{'i', 'j', 'k', 'l'},
					{'m', 'n', 'o', 'p'},
				},
				word: "abcd",
			},
			want: 0,
		},
		{
			name: "Test2",
			args: args{
				puzzle: [][]byte{
					{'a', 'f', 'a'},
					{'x', 'o', 'b'},
					{'o', 'r', 'c'},
				},
				word: "abc",
			},
			want: 0,
		},
		{
			name: "Test3",
			args: args{
				puzzle: [][]byte{
					{'z', 'e', 'z'},
					{'w', 'n', 'a'},
					{'n', 'q', 'b'},
				},
				word: "abc",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find_the_word(tt.args.puzzle, tt.args.word); got != tt.want {
				t.Errorf("find_the_word() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotate_clockwise(t *testing.T) {
	type args struct {
		puzzle [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Test4-rotate",
			args: args{
				puzzle: [][]byte{
					{'z', 'e', 'z'},
					{'w', 'n', 'a'},
					{'n', 'q', 'b'},
				},
			},
			want: [][]byte{
				{'a', 'f', 'a'},
				{'x', 'o', 'b'},
				{'o', 'r', 'c'},
			},
		},
		{
			name: "Test5-rotate",
			args: args{
				puzzle: [][]byte{
					{'a', 'b', 'c'},
					{'d', 'e', 'f'},
					{'g', 'h', 'i'},
				},
			},
			want: [][]byte{
				{'b', 'c', 'd'},
				{'e', 'f', 'g'},
				{'h', 'i', 'j'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate_clockwise(tt.args.puzzle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate_clockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVertical_search(t *testing.T) {
	type args struct {
		puzzle [][]byte
		word   string
	}
	tests := []struct {
		name string
		args args
		want bool
		x    int
		y    int
	}{
		{
			name: "Test6-vertical",
			args: args{
				puzzle: [][]byte{
					{'a', 'b', 'c'},
					{'d', 'e', 'f'},
					{'g', 'h', 'i'},
				},
				word: "abc",
			},
			want: false,
			x:    0,
			y:    0,
		},
		{
			name: "Test7-vertical",
			args: args{
				puzzle: [][]byte{
					{'z', 'e', 'a'},
					{'w', 'n', 'b'},
					{'n', 'q', 'c'},
				},
				word: "abc",
			},
			want: true,
			x:    2,
			y:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, x, y := vertical_search(tt.args.puzzle, tt.args.word); got == tt.want && x == tt.x && y == tt.y {
				fmt.Printf("%v %v %v\n", got, x, y)
			} else {
				t.Errorf("vertical_search(): got = %v, want %v | x_got = %v, x_actual = %v | y_got = %v, y_actual = %v ", got, tt.want, x, tt.x, y, tt.y)
			}
		})
	}
}

func TestHorizontal_search(t *testing.T) {
	type args struct {
		puzzle [][]byte
		word   string
	}
	tests := []struct {
		name string
		args args
		want bool
		x    int
		y    int
	}{
		{
			name: "Test8-horizontal",
			args: args{
				puzzle: [][]byte{
					{'a', 'b', 'c'},
					{'d', 'e', 'f'},
					{'g', 'h', 'i'},
				},
				word: "abc",
			},
			want: true,
			x:    0,
			y:    0,
		},
		{
			name: "Test9-horizontal",
			args: args{
				puzzle: [][]byte{
					{'a', 'd', 'g'},
					{'b', 'e', 'f'},
					{'c', 'h', 'i'},
				},
				word: "abc",
			},
			want: false,
			x:    0,
			y:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, x, y := horizontal_search(tt.args.puzzle, tt.args.word); got == tt.want && x == tt.x && y == tt.y {
				fmt.Printf("%v %v %v\n", got, x, y)
			} else {
				t.Errorf("horizontal_search(): got = %v, want %v | x_got = %v, x_actual = %v | y_got = %v, y_actual = %v ", got, tt.want, x, tt.x, y, tt.y)
			}
		})
	}
}
