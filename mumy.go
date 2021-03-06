package main

import (
	"strings"
)

var vowels = [10] string {
	"a",
	"e",
	"i",
	"o",
	"u",
	"A",
	"E",
	"I",
	"O",
	"U",
}

func IsVowel(letter string) bool {
	for _, vowel := range vowels {
		if vowel == letter {
			return true
		}
	}
	return false
}


func CalculatePercentage(word string) float64 {

	var totalLetters int  = len(word)

	var totalVowels int = 0
	var letters = strings.Split(word,"")
	for _, letter := range letters{
		if IsVowel(letter){
			totalVowels++
		}
	}

	var percentage float64 = float64(totalVowels) / float64(totalLetters)
	return percentage
}

func ConvertVowelsToMumy(word string) string{
	if CalculatePercentage(word) <= 0.3  {
		return word
	}

	var mumyword string = word
	var letters = strings.Split(mumyword,"")
	for _, letter := range letters{
		if IsVowel(letter){
			strings.Replace(mumyword, letter,"mumy",-1)
		}git
	}

	return mumyword
}
