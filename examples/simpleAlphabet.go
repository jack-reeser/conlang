package main

import (
	"fmt"
	"strings"

	"github.com/jack-reeser/conlang/alphabet"
	"github.com/jack-reeser/conlang/common"
)

const (
	COMMA     = ","
	CONSONANT = alphabet.Class('C')
	VOWEL     = alphabet.Class('V')
)

func main() {
	fmt.Println("An example of a simple alphabet...")

	// parse in a simple alphabet
	inputLetters := map[alphabet.Class][]string{
		VOWEL: {
			"A,a",
			"E,e",
			"I,i",
			"O,o",
			"U,u",
		},
		CONSONANT: {
			"B,b",
			"C,c",
			"D,d",
			"F,f",
			"G,g",
			"H,h",
			"J,j",
			"K,k",
			"L,l",
			"M,m",
			"N,n",
			"P,p",
			"Q,q",
			"R,r",
			"S,s",
			"T,t",
			"V,v",
			"W,w",
			"X,x",
			"Z,z",
		},
	}

	letters := []alphabet.Letter{}

	addLetter := func(class alphabet.Class, csvList []string) {
		for _, csv := range csvList {
			if values := strings.Split(csv, COMMA); len(values) == 2 {
				letters = append(
					letters,
					alphabet.NewLetter(values[0], values[1], class),
				)
			}
		}
	}

	for class, list := range inputLetters {
		addLetter(class, list)
	}

	fmt.Printf("Parsed an alphabet of length %d\n", len(letters))

	simpleAlphabet := alphabet.New(letters)

	classMap := map[alphabet.Class]common.Collection[alphabet.Letter]{
		CONSONANT: simpleAlphabet.GetLettersByClass(CONSONANT),
		VOWEL:     simpleAlphabet.GetLettersByClass(VOWEL),
	}

	fmt.Printf("Got %d consonants and %d vowels\n", classMap[CONSONANT].Len(), classMap[VOWEL].Len())

	// make a function to get a random word using a pattern
	getRandomWord := func(pattern string) string {
		var randomWord string
		for i, class := range alphabet.StringToClasses(pattern) {
			var letter alphabet.Letter
			if list, ok := classMap[class]; ok {
				letter = list.GetRandom()
			} else {
				letter = alphabet.NewLetter("?", "?", '?')
			}
			if i == 0 {
				randomWord = letter.Upper()
			} else {
				randomWord = randomWord + letter.Lower()
			}
		}
		return randomWord
	}

	fmt.Println("Generated random words:")

	for _, pattern := range []string{"CVC", "CVCV", "VC", "V", "VCV", "#"} {
		fmt.Printf("%s ", getRandomWord(pattern))
	}
}
