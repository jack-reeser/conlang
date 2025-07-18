package morph

import "fmt"

// Lexemes represent words, more or less. A Lexeme is generally subject to
// rules of inflection and rules of word formation.
type Lexeme interface {
	fmt.Stringer
	// Index returns the lowercased dictionary index of the lexeme.
	Index() string
	// Class returns the morphological class
	Class() Class
}
