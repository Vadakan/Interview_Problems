package main

import (
	"fmt"
	"strings"
)

var a string

func main() {

	a = "SUndAr"

	count_of_a := strings.Count(a, "a")
	count_of_e := strings.Count(a, "e")
	count_of_i := strings.Count(a, "i")
	count_of_o := strings.Count(a, "o")
	count_of_u := strings.Count(a, "u")

	total := count_of_a + count_of_e + count_of_i + count_of_o + count_of_u

	fmt.Println("number of Vowels :", total)
	fmt.Println("number of consonants :", len(a)-total)

}

