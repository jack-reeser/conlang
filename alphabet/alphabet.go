package alphabet

import "github.com/jack-reeser/conlang/common"

// Alphabet represents a collection of Letters
type Alphabet interface {
	// GetLetters returns all Letters
	GetLetters() []Letter
	// GetLettersByClass returns all Letters that match the given Class
	GetLettersByClass(Class) []Letter
	// GetClasses returns a slice of all unique Classes
	GetClasses() []Class
}

// New takes a list of Letters and returns an Alphabet
func New(letters []Letter) Alphabet {
	return basicAlphabet{common.CollectionFrom[Letter](letters)}
}

type basicAlphabet struct {
	common.Collection[Letter]
}

func (b basicAlphabet) GetLetters() []Letter { return b.ToSlice() }
func (b basicAlphabet) GetLettersByClass(c Class) []Letter {
	return b.Select(
		func(l Letter) bool {
			return l.IsClass(c)
		})
}
func (b basicAlphabet) GetClasses() []Class {
	alphabetClassSet := map[Class]bool{}
	for _, letter := range b.ToSlice() {
		for class := range letter.GetClassMap() {
			alphabetClassSet[class] = true
		}
	}
	return common.CollectionFrom[Class](alphabetClassSet).ToSlice()
}
