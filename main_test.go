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
		{
			name: "no word found",
			args: args{
				puzzle: [][]byte{
					{'k', 'w', 'h', 'd', 'e', 't', 'a', 'w', 's', 'u', 'w', 's', 'i', 'b', 'o', 'v', 'z', 'e', 'z', 'b'},
					{'l', 'v', 'd', 'd', 'c', 'j', 'z', 'w', 'o', 'd', 'g', 'v', 'e', 'y', 's', 'e', 'k', 'w', 'w', 'm'},
					{'o', 'b', 'w', 'j', 'e', 'k', 'y', 'u', 'm', 'j', 'm', 'a', 'f', 'i', 's', 'w', 'o', 'r', 'u', 'c'},
					{'q', 'l', 'h', 'y', 'x', 'c', 'p', 'y', 'e', 'q', 'p', 'y', 't', 't', 'a', 'x', 'u', 'x', 'k', 'b'},
					{'e', 'x', 'v', 'd', 'h', 'z', 'v', 'z', 'n', 'h', 'q', 'e', 'c', 'i', 'h', 'd', 'l', 'x', 'c', 'k'},
					{'m', 'w', 't', 'o', 'e', 'c', 'm', 'h', 'g', 'k', 'b', 'k', 'x', 't', 'e', 't', 'h', 'a', 'p', 'y'},
					{'j', 'o', 'y', 'b', 'e', 'y', 'i', 'a', 'y', 'j', 'n', 'y', 'w', 'r', 'q', 'x', 'v', 'p', 'o', 'h'},
					{'e', 'p', 'r', 'w', 'm', 'p', 'b', 't', 'z', 'i', 'c', 'c', 'x', 'l', 'g', 'm', 'f', 'z', 'i', 't'},
					{'s', 'e', 'i', 'a', 'm', 'o', 'q', 'p', 'v', 'y', 'x', 'r', 's', 'c', 'r', 'w', 't', 'g', 'm', 'f'},
					{'z', 'o', 'z', 'i', 'y', 'v', 'e', 'p', 'y', 'y', 'w', 'a', 'i', 'y', 'p', 't', 'm', 'e', 's', 'd'},
					{'p', 's', 'c', 'i', 'p', 'q', 'w', 'g', 'l', 'k', 'd', 'o', 'w', 'y', 's', 'o', 'b', 'v', 'j', 'm'},
					{'t', 'o', 'u', 'k', 'i', 'j', 'h', 'j', 'q', 'l', 'l', 'y', 'z', 'c', 'u', 'a', 'l', 'b', 'q', 'g'},
					{'x', 'y', 'k', 'd', 'r', 'b', 'x', 'z', 'v', 'o', 'j', 'd', 's', 'k', 'x', 'e', 'q', 's', 'p', 'c'},
					{'p', 't', 'e', 'x', 's', 'a', 'm', 'm', 'l', 'i', 'c', 'n', 'z', 'q', 'j', 'm', 'y', 'y', 'k', 'd'},
					{'s', 's', 'n', 't', 'w', 'n', 'i', 'd', 'c', 'm', 'x', 'l', 'f', 'j', 'g', 'v', 'm', 'z', 'y', 'v'},
					{'e', 'j', 'e', 'n', 'm', 'f', 'w', 'l', 's', 'h', 'h', 'c', 'm', 'z', 'u', 't', 'x', 'q', 'j', 'c'},
					{'x', 'f', 's', 's', 'h', 'h', 'a', 'p', 'x', 'q', 'y', 'v', 'l', 'm', 'h', 'q', 'n', 's', 's', 'k'},
					{'y', 'x', 'q', 'n', 'i', 'v', 'v', 'd', 'k', 'b', 'b', 'c', 'e', 'e', 'y', 'd', 'r', 'a', 'n', 'b'},
					{'l', 'f', 'i', 's', 'g', 'p', 'p', 'l', 'm', 'a', 'z', 'n', 'm', 'd', 'z', 'w', 'k', 'h', 's', 'b'},
					{'n', 'j', 'q', 'e', 'g', 'e', 'm', 'a', 'j', 'm', 'i', 'z', 'x', 'i', 'o', 'g', 'x', 'i', 'x', 'i'},
				},
				word: "netapp",
			},
			want: 1000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := find_the_word(tt.args.puzzle, tt.args.word)
			if got != tt.want {
				t.Errorf("find_the_word() = %v, want %v", got, tt.want)
			}
			if got == 1000001 {
				fmt.Println("No word found")
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
