package palindrome

import "strings"

//palindrome functions using looping
//parameter string
//return boolean
func PalindromeLoop(value string) bool {
	for i := 0; i < len(value); i++ {
		str := strings.Split(value, "")
		var nilaiAwal = str[i]
		var nilaiAkhir = str[len(value)-i-1]
		if nilaiAwal != nilaiAkhir {
			return false
		}
	}
	return true
}
