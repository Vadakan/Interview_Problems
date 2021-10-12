package main

import "fmt"

var count_v int
var count_c int

func count_Vowels(inp string) (Vowels int, Consonants int) {

	var dump = make(map[string]string)

	dump["A"] = "A"
	dump["E"] = "E"
	dump["I"] = "I"
	dump["O"] = "O"
	dump["U"] = "U"
	dump["a"] = "a"
	dump["e"] = "e"
	dump["i"] = "i"
	dump["o"] = "o"
	dump["u"] = "u"

	for _, value := range inp {
		temp := string(value)
		if _, ok := dump[temp]; ok {
			count_v++
		} else {
			count_c++
		}

	}
	return count_v, count_c

}

func main() {

	count_Vowels, count_Consonants := count_Vowels("SUNDAR")
	fmt.Println("count of Vowels :", count_Vowels)
	fmt.Println("count of Consonants :", count_Consonants)

}

