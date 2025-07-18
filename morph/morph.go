package morph

import "fmt"

// Morphemes represent a broad category of morphological elements, ranging from
// affixes to word roots. Morphemes may be classified as "free" or "bound". A
// "free" morpheme is independent of any other morpheme, and may appear within
// lexemes. "Bound" morphemes only appear as dependent parts of words, and may
// be bound to other bound morphemes.
type Morpheme interface {
	fmt.Stringer
	// IsFree indicates if the Morpheme is free. If it is not free, it must be bound.
	IsFree() bool
	// IsPrefix indicates if the Morpheme precedes or follows other Morphemes.
	IsPrefix() bool
	// Combine combines two morphemes together. If they are both free morphemes,
	// then the receiver morpheme will be ordered first in the output. If one
	// is free and the other is bound, the position of the bound morpheme will
	// dictate the position it appears in the output. If both morphemes are bound
	// and a.IsPrefix() != b.IsPrefix(), then the output will be ordered according
	// to a and b's respective prefix/suffix ordering. If both morphemes are bound
	// and a.IsPrefix() == b.IsPrefix(), then the receiver morpheme will be ordered
	// first in the output.
	Combine(Morpheme) Morpheme
}

// NewPrefix makes a new prefix Morpheme
func NewPrefix(s string) Morpheme {
	return boundMorpheme{s, true}
}

// NewStem makes a new stem Morpheme
func NewStem(s string) Morpheme {
	return freeMorpheme(s)
}

// NewSuffix makes a new suffix Morpheme
func NewSuffix(s string) Morpheme {
	return boundMorpheme{s, false}
}

// NewMorpheme makes a new Morpheme.
func NewMorpheme(s string, free, prefix bool) Morpheme {
	if free {
		return freeMorpheme(s)
	}
	return boundMorpheme{s, prefix}
}

type boundMorpheme struct {
	morpheme string
	prefix   bool
}

func (b boundMorpheme) IsFree() bool   { return false }
func (b boundMorpheme) IsPrefix() bool { return b.prefix }
func (b boundMorpheme) String() string { return b.morpheme }
func (b boundMorpheme) Combine(other Morpheme) Morpheme {
	if other.IsFree() {
		/*
		 if a bound morpheme combines with a free one, we put it in place
		 and call the new morpheme free.
		 example: ("-ly", "cool") => "coolly"

		 [note: this is a weird example that needs to be handled by complicated
		 english orthography rules after the two morphemes are combined; that
		 is, after the production of an adverb, endings of "lly" and "ely" are
		 replaced by "ly".]
		*/
		if b.IsPrefix() {
			return NewMorpheme(b.String()+other.String(), true, true)
		}
		return NewMorpheme(other.String()+b.String(), true, true)
	} else if b.IsPrefix() != other.IsPrefix() {
		/*
		 If two bound morphemes are prefix + suffix, we combine them in that
		 order and call the new morpheme free. this probably shouldn't happen
		 in some cases, but this is too low level to enforce, so it's allowed.
		 example: ("un-", "-ly") => "unly"

		 [note: morphological rules that require a root/stem morpheme must be
		 handled at a higher level. "unly" doesn't and probably can't mean
		 anything significant in english because word formation rules would
		 require a root or stem to produce a meaning such as "unruly".]
		*/
		if b.IsPrefix() {
			return NewMorpheme(b.String()+other.String(), true, true)
		}
		return NewMorpheme(other.String()+b.String(), true, true)
	}
	/*
	 By process of elimination, we have arrived at a case where both morphemes
	 are bound, and have the same IsPrefix value. In this instance, the
	 receiver morpheme is ordered first, and we call the output morpheme bound
	 with the IsPrefix value that both morphemes have.
	 example: ("un-", "re-") => "unre-"

	 [note: rules that enforce strict ordering of certain classes of morphemes
	 must be handled at a higher level than this. given this example, consider
	 that the only instances of words beginning with "reun-" regard "re-union"
	 or "re-unite". in these words, "un-" is not the "un-" you see in "undo".]
	*/
	return NewMorpheme(b.String()+other.String(), false, b.IsPrefix())
}

type freeMorpheme string

func (f freeMorpheme) IsFree() bool   { return true }
func (f freeMorpheme) IsPrefix() bool { return false }
func (f freeMorpheme) String() string { return string(f) }
func (f freeMorpheme) Combine(other Morpheme) Morpheme {
	// example: ("dog", "house") => "doghouse"
	if other.IsFree() {
		return NewMorpheme(f.String()+other.String(), true, true)
	}
	// example: ("do", "re-") => "redo"
	if other.IsPrefix() {
		return NewMorpheme(other.String()+f.String(), true, true)
	}
	// example: ("do", "-ing") => "doing"
	return NewMorpheme(f.String()+other.String(), true, true)
}
