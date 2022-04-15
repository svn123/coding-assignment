package main

import (
	"fmt"
)

//Map used for rotation
var rotation_map = map[byte]byte{
	'a': 'b',
	'b': 'c',
	'c': 'd',
	'd': 'e',
	'e': 'f',
	'f': 'g',
	'g': 'h',
	'h': 'i',
	'i': 'j',
	'j': 'k',
	'k': 'l',
	'l': 'm',
	'm': 'n',
	'n': 'o',
	'o': 'p',
	'p': 'q',
	'q': 'r',
	'r': 's',
	's': 't',
	't': 'u',
	'u': 'v',
	'v': 'w',
	'w': 'x',
	'x': 'y',
	'y': 'z',
	'z': 'a',
}

var count int = 0

// This will rotate the characters present in the 2d array if the string is not found in the present run
func rotate_clockwise(puzzle [][]byte) [][]byte {
	var new_puzzle [][]byte
	for i := 0; i < len(puzzle); i++ {
		var new_row []byte
		for j := 0; j < len(puzzle[i]); j++ {
			new_row = append(new_row, rotation_map[puzzle[i][j]])
		}
		new_puzzle = append(new_puzzle, new_row)
	}
	fmt.Println(new_puzzle)
	return new_puzzle
}

// search character in vertical direction
func vertical_search(puzzle [][]byte, word string) bool {
	length := len(word)
	var str string
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[j][i] == word[0] {
				str = ""
				for k := 0; k < length; k++ {
					if j+k < len(puzzle[i]) {
						str += string(puzzle[j+k][i])
					}
					fmt.Println(str)
				}
				if str == word {
					fmt.Println("The co-ordinates are : ", i, j)
					return true
				}
			}
		}
	}
	return false
}

// search character in horizontal direction
func horizontal_search(puzzle [][]byte, word string) bool {
	length := len(word)
	var str string
	for i := 0; i < len(puzzle); i++ {
		for j := i; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == word[0] {
				str = ""
				for k := 0; k < length; k++ {
					if j+k < len(puzzle[i]) {
						str += string(puzzle[i][j+k])
					}
					fmt.Println(str)
				}

				if str == word {
					fmt.Println("The co-ordinates are : ", i, j)
					return true
				}
			}
		}
	}
	return false
}

// function for searching for a particular word in the 2d array
func find_the_word(puzzle [][]byte, word string) int {
	var puzzle_rotated [][]byte
	if horizontal_search(puzzle, word) || vertical_search(puzzle, word) {
		print(count)
	} else {
		puzzle_rotated = rotate_clockwise(puzzle)
		count++
		find_the_word(puzzle_rotated, word)
	}
	return count
}

func main() {

	sample := [][]byte{{'z', 'e', 'z'},
		{'w', 'n', 'a'},
		{'n', 'q', 'b'}}

	/*
		sample2 := [][]byte{{'a', 'f', 'a'},
			{'x', 'o', 'b'},
			{'o', 'r', 'c'}}
	*/
	/*
		puzzle1 := [][]byte{{'k', 'w', 'h', 'd', 'e', 't', 'a', 'w', 's', 'u', 'w', 's', 'i', 'b', 'o', 'v', 'z', 'e', 'z', 'b'},
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
			{'n', 'j', 'q', 'e', 'g', 'e', 'm', 'a', 'j', 'm', 'i', 'z', 'x', 'i', 'o', 'g', 'x', 'i', 'x', 'i'}}

			puzzle2 := [][]byte{{'i', 'w', 'j', 'm', 'j', 'w', 'v', 'v', 't', 'x', 'm', 'p', 'x', 'b', 's', 'j', 'k', 'i', 'f', 'r'},
				{'q', 'p', 'k', 'l', 'j', 'x', 'a', 'o', 'd', 'b', 'd', 'v', 's', 'y', 'k', 't', 'l', 'd', 'w', 'n'},
				{'g', 'y', 'b', 'x', 'j', 'p', 'd', 'q', 'v', 'p', 'l', 'j', 'j', 'b', 'p', 'd', 'n', 'p', 'y', 'k'},
				{'s', 'v', 'n', 'y', 'h', 'r', 't', 'c', 'o', 'w', 'x', 'd', 'c', 'o', 'y', 's', 'k', 'x', 'q', 'j'},
				{'w', 'v', 'e', 'u', 'a', 'd', 'b', 'x', 's', 'j', 'w', 'u', 'y', 'f', 'x', 'j', 'k', 'b', 'd', 'g'},
				{'c', 'q', 's', 'y', 's', 'b', 'q', 'e', 'l', 'i', 'c', 's', 'z', 'e', 'w', 'p', 'm', 'c', 'z', 'z'},
				{'z', 'a', 'h', 'd', 'e', 'h', 'w', 'd', 'q', 'v', 'j', 'd', 'z', 'b', 'f', 'y', 'w', 'i', 'z', 'v'},
				{'o', 'o', 'q', 'g', 'l', 'g', 'k', 'h', 'h', 'l', 'r', 'y', 'w', 'n', 's', 'p', 'h', 'g', 's', 'w'},
				{'i', 'j', 'l', 'g', 'b', 'i', 'n', 'c', 'e', 'x', 'w', 'v', 'z', 'g', 'h', 'c', 'b', 'z', 'd', 'v'},
				{'e', 'e', 'k', 'q', 'v', 'y', 'n', 'k', 'a', 'r', 'n', 'k', 'i', 'v', 'n', 'j', 'z', 't', 'c', 'u'},
				{'y', 'z', 'k', 's', 'y', 'n', 'p', 't', 'p', 'l', 'a', 'q', 'i', 'd', 'a', 'h', 'z', 's', 'a', 's'},
				{'g', 'n', 'r', 'o', 'f', 'e', 'g', 'b', 'u', 'b', 'b', 'v', 'a', 't', 'k', 'g', 'z', 'l', 'w', 'm'},
				{'d', 'v', 'f', 'i', 'v', 'g', 'f', 'j', 's', 't', 'e', 'i', 'l', 'q', 's', 'f', 's', 'm', 'z', 'n'},
				{'q', 'c', 'u', 'e', 'z', 'l', 'n', 'i', 'n', 't', 'x', 'o', 'p', 'h', 's', 'n', 't', 'e', 'f', 'k'},
				{'u', 'e', 'u', 'r', 'u', 'f', 'n', 'y', 'r', 'y', 'l', 'i', 'v', 'm', 'l', 'p', 'n', 'c', 'g', 'e'},
				{'g', 'o', 'j', 'q', 'i', 'q', 'l', 'm', 'l', 'g', 'a', 'e', 'x', 'i', 'q', 'g', 'n', 'z', 'a', 'b'},
				{'h', 'm', 'd', 'w', 'c', 'u', 'h', 'm', 'q', 'p', 'r', 'd', 's', 'q', 'n', 'b', 'y', 'k', 'v', 'c'},
				{'x', 'o', 'm', 'i', 'z', 'r', 'w', 'b', 'v', 'd', 't', 'g', 'k', 'm', 'r', 't', 'p', 'v', 'n', 's'},
				{'y', 'y', 'x', 'i', 'u', 'd', 'k', 'e', 't', 'o', 'u', 'v', 'a', 's', 'c', 'z', 's', 'x', 'j', 'x'},
				{'j', 'e', 'x', 'f', 'k', 'l', 'h', 'l', 'r', 'p', 's', 'b', 'q', 'r', 'c', 'h', 'x', 'l', 'x', 'd'}}
	*/
	//puz1_srchterm := "netapp"
	//puz2_srchterm := "containers"
	sample_word := "abc"

	find_the_word(sample, sample_word)
	//find_the_word(puzzle1, puz1_srchterm)
	//find_the_word(puzzle2, puz2_srchterm)
	//horizontal_search(puzzle1, puz1_srchterm)
	//horizontal_search(puzzle2, puz2_srchterm)
	//vertical_search(puzzle1, puz1_srchterm)
	//vertical_search(puzzle2, puz2_srchterm)

	//vertical_search(sample2, sample_word)

}
