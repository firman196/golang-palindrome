package palindrome

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsPalindrome(t *testing.T) {
	//testing palindrome loop
	result := PalindromeLoop("katak")
	assert.Equal(t, true, result)
}

func TestIsNotPalindrome(t *testing.T) {
	//testing palindrome loop
	result := PalindromeLoop("Test")
	assert.Equal(t, false, result)
}
