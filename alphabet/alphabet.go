package alphabet

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
	return basicAlphabet(letters)
}

type basicAlphabet []Letter

func (b basicAlphabet) GetLetters() []Letter { return b }
func (b basicAlphabet) GetLettersByClass(c Class) []Letter {
	found := make([]Letter, 0)
	for _, letter := range b {
		if letter.IsClass(c) {
			found = append(found, letter)
		}
	}
	return found
}
func (b basicAlphabet) GetClasses() []Class {
	alphabetClassSet := map[Class]bool{}
	for _, letter := range b {
		for class := range letter.GetClassMap() {
			alphabetClassSet[class] = true
		}
	}
	allClasses := make([]Class, 0)
	for value := range alphabetClassSet {
		allClasses = append(allClasses, value)
	}
	return allClasses
}
