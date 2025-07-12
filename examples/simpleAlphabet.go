package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/jack-reeser/conlang/alphabet"
)

func main() {
	fmt.Println("An example of a simple alphabet...")

	// parse in a simple alphabet
	inputLetters := []string{
		"A,a,V",
		"E,e,V",
		"I,i,V",
		"O,o,V",
		"U,u,V",
		"B,b,C",
		"C,c,C",
		"D,d,C",
		"F,f,C",
		"G,g,C",
		"H,h,C",
		"J,j,C",
		"K,k,C",
		"L,l,C",
		"M,m,C",
		"N,n,C",
		"P,p,C",
		"Q,q,C",
		"R,r,C",
		"S,s,C",
		"T,t,C",
		"V,v,C",
		"W,w,C",
		"X,x,C",
		"Z,z,C",
	}

	letters := []alphabet.Letter{}
	const COMMA string = ","

	for _, letterCsv := range inputLetters {
		values := strings.Split(letterCsv, COMMA)
		if len(values) != 3 {
			continue
		} else {
			letter := alphabet.NewLetter(
				values[0], values[1], alphabet.StringToClasses(values[2])...)

			letters = append(letters, letter)
		}
	}

	fmt.Printf("Parsed an alphabet of length %d\n", len(letters))

	// separate the letters by class
	const C = alphabet.Class('C')
	const V = alphabet.Class('V')

	getLettersByClass := func(c alphabet.Class) []alphabet.Letter {
		foundList := make([]alphabet.Letter, 0)
		for _, letter := range letters {
			if letter.IsClass(c) {
				foundList = append(foundList, letter)
			}
		}
		return foundList
	}

	consonants := getLettersByClass(C)
	vowels := getLettersByClass(V)

	fmt.Printf("Got %d consonants and %d vowels\n", len(consonants), len(vowels))

	// make a function to get a random letter
	getRandomLetter := func(c alphabet.Class) alphabet.Letter {
		list := map[alphabet.Class][]alphabet.Letter{
			C: consonants,
			V: vowels,
		}[c]
		if len(list) > 0 {
			i := rand.Intn(len(list))
			return list[i]
		}
		return alphabet.NewLetter("?", "?", alphabet.Class('?'))
	}

	// make a function to get a random word using a pattern
	getRandomWord := func(pattern string) string {
		var randomWord string
		for i, class := range alphabet.StringToClasses(pattern) {
			letter := getRandomLetter(class)
			if i == 0 {
				randomWord = letter.Upper()
			} else {
				randomWord = randomWord + letter.Lower()
			}
		}
		return randomWord
	}

	fmt.Println("Generated random words")

	for _, pattern := range []string{"CVC", "CVCV", "VC", "V", "VCV", "#"} {
		fmt.Printf("%s ", getRandomWord(pattern))
	}
}
