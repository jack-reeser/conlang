package alphabet

// Letter is a symbol that represents a sound.
type Letter interface {
	// Upper returns a string that represents the uppercase version of the letter.
	Upper() string
	// Lower returns a string that represents the lowercase version of the letter.
	Lower() string
	// IsClass returns true if this Letter matches the given Class.
	IsClass(Class) bool
	// GetClasses returns a list of all Classes that typify the Letter.
	GetClassSlice() []Class
	// GetClassMap returns the letter's underlying Class map.
	GetClassMap() map[Class]bool
}

// NewLetter makes a new letter given an upper, lower, and variable number of
// Letter Classes.
func NewLetter(upper, lower string, classes ...Class) Letter {
	letterClassSet := map[Class]bool{}
	for _, class := range classes {
		letterClassSet[class] = true
	}

	if upper == lower {
		return simpleLetter{
			letter:   lower,
			classSet: letterClassSet,
		}
	}

	return fullLetter{
		upper:    upper,
		lower:    lower,
		classSet: letterClassSet,
	}
}

// classSet is a reusable private set of Classes
type classSet map[Class]bool

func (c classSet) IsClass(class Class) bool    { return c[class] }
func (c classSet) GetClassMap() map[Class]bool { return c }
func (c classSet) GetClassSlice() []Class {
	classes := make([]Class, len(c))
	i := 0
	for class := range c {
		classes[i] = class
		i++
	}
	return classes
}

// simpleLetter represents letters that do not have distinct upper and lower values
type simpleLetter struct {
	letter string
	classSet
}

func (s simpleLetter) Upper() string { return s.letter }
func (s simpleLetter) Lower() string { return s.letter }

// fullLetter represents letters with distinct upper and lower values
type fullLetter struct {
	upper string
	lower string
	classSet
}

func (s fullLetter) Upper() string { return s.upper }
func (s fullLetter) Lower() string { return s.lower }
