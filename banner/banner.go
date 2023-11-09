package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// BANNER
	// banner("Go", 6)
	// banner("G😓", 6)
	// s := "G😓"
	// fmt.Println("len:", len(s))
	// for i, r := range s {
	// 	fmt.Println(i, r)
	// 	if i == 0 {
	// 		fmt.Printf("%c of type %T\n", r, r)
	// 		// returns G of type int32
	// 		// rune
	// 	}
	// }
	// // byte (uint8)
	// // rune (int32)
	// b := s[0]
	// fmt.Printf("%c of type %T\n", b, b)
	// // returns G of type uint8
	// // byte uint8

	// x, y := 1, "1"
	// fmt.Printf("x=%v, y=%v\n", x, y)
	// fmt.Printf("x=%#v, y=%#v\n", x, y) // Use #v in debug/log
	// fmt.Printf("%20s!", s)

	// IS PALINDROME
	// res := isPalindrome("g")
	// fmt.Println("is g a palindrome?")
	// fmt.Println(res)
	// res = isPalindrome("go")
	// fmt.Println("is go a palindrome?")
	// fmt.Println(res)
	// res = isPalindrome("gog")
	// fmt.Println("is gog a palindrome?")
	// fmt.Println(res)
	// res = isPalindrome("gogo")
	// fmt.Println("is gogo a palindrome?")
	// fmt.Println(res)

}

func isPalindrome(s string) bool {
	// rs := []rune(s) // a slice of runes out of s -> unicode safe
	// for i := 0; i < len(rs)/2; i++ {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2 // BUG: len is in bytes
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()

}
