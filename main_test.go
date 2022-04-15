package main

import "testing"

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

func TestFind_the_word(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find_the_word(tt.args.puzzle, tt.args.word); got != tt.want {
				t.Errorf("find_the_word() = %v, want %v", got, tt.want)
			}
		})
	}
}
