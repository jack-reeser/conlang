package alphabet

// Class is a rune that represents a distinct class of letters.
type Class rune

// StringToClasses converts a string into a slice of Classes. If the given string
// has a length of zero, an empty slice is returned.
func StringToClasses(s string) []Class {
	if len(s) <= 0 {
		return []Class{}
	}

	classes := make([]Class, len(s))
	for i, char := range s {
		classes[i] = Class(char)
	}

	return classes
}
