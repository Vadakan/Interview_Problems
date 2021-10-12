package main

import (
	"fmt"
)

var Vowel_ctr, cons_ctr int

func IsVowel(inp string) (V_Count int, C_count int) {
	for _, char := range inp {
		switch temp := string(char); temp {
		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
			Vowel_ctr++
		case "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z":
			cons_ctr++
		default:
			continue
		}

	}
	return Vowel_ctr, cons_ctr
}

func main() {

	Vowel_ctr, cons_ctr := IsVowel("SundAr12345__++++")
	fmt.Println("Number of Vowels :", Vowel_ctr)
	fmt.Println("Number of Consonants :", cons_ctr)

}

