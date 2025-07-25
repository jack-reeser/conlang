package common

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

// Collection is a generic collection of comparable values.
type Collection[T comparable] interface {
	fmt.Stringer
	// Select returns a List of values that satisfy the given condition.
	Select(func(T) bool) List[T]
	// ToSlice returns a slice of comparable values.
	ToSlice() []T
	// ToMap returns a set map of comparable values.
	ToMap() map[T]bool
	// ToList returns a List of comparable values.
	ToList() List[T]
	// ToSet returns a Set of comparable values.
	ToSet() Set[T]
	// GetRandom returns a single randomly chosen value from the Collection. If
	// the Collection is empty, a null value of type T is returned.
	GetRandom() T
	// ToShuffledList returns a random permutation of the Collection as List[T].
	ToShuffledList() List[T]
	// ToSortedList takes a sort function and returns a sorted List[T].
	ToSortedList(func(T, T) int) List[T]
	// Len returns the size of the Collection
	Len() int
}

// CollectionFrom creates a new Collection of comparable types. This function
// chooses the appropriate Collection implementation based on the type given.
// Base implementations such as List[T] and Set[T] may be passed to this
// function to convert them to a Collection[T].
//
// If an appropriate Collection implementation cannot be inferred from the given
// source type, an empty Set[T] is returned.
func CollectionFrom[T comparable](a any) (collection Collection[T]) {
	switch t := a.(type) {
	case map[T]bool:
		collection = Set[T](t)
	case []T:
		collection = List[T](t)
	case List[T]:
		collection = t
	case Set[T]:
		collection = t
	default:
		collection = Set[T]{}
	}
	return
}

// List is generic slice of comparable types.
type List[T comparable] []T

// Select returns a new List of items that satisfy the given condition.
func (l List[T]) Select(f func(T) bool) List[T] {
	selected := make([]T, 0)
	for _, item := range l {
		if f(item) {
			selected = append(selected, item)
		}
	}
	return selected
}

// ToSlice returns the underlying generic slice.
func (l List[T]) ToSlice() []T { return l }

// ToMap converst the List to a Set, then returns the underlying Set's map.
func (l List[T]) ToMap() map[T]bool { return l.ToSet().ToMap() }

// ToList returns the List as a list. This is not useful and only exists to
// satisfy Collection[T].
func (l List[T]) ToList() List[T] { return l }

// ToSet converts the list to a Set.
func (l List[T]) ToSet() Set[T] {
	if len(l) == 0 {
		return Set[T]{}
	}

	set := make(Set[T])
	for _, item := range l {
		set[item] = true
	}
	return set
}

// GetRandom efficiently returns a random type T from the List.
func (l List[T]) GetRandom() (chosen T) {
	if len(l) == 0 {
		return
	} else {
		chosen = l[rand.Intn(len(l))]
	}
	return
}

// ToShuffledList returns the List[T] back with its values shuffled.
func (l List[T]) ToShuffledList() (shuffled List[T]) {
	size := len(l)
	if size == 0 {
		shuffled = l
	} else {
		shuffled = make([]T, size)
		for i, value := range rand.Perm(size) {
			shuffled[i] = l[value]
		}
	}
	return
}

// ToSortedList copies the original List and sorts it according to the given sort
// function. The sort function should return 1 when a > b, -1 when a < b, and 0
// when a == b.
func (l List[T]) ToSortedList(f func(T, T) int) (sorted List[T]) {
	if len(l) <= 1 {
		sorted = l
	} else {
		sorted = make([]T, len(l))
		copy(sorted, l)
		slices.SortFunc(sorted, f)
	}
	return
}

// Len returns the length of the underlying List slice.
func (l List[T]) Len() int { return len(l) }

func (l List[T]) String() string {
	switch v := any(l).(type) {
	case []string:
		return strings.Join(v, "")
	case []rune:
		return string(v)
	case []fmt.Stringer:
		s := make([]string, len(v))
		for i := 0; i < len(v); i++ {
			s[i] = v[i].String()
		}
		return strings.Join(s, "")
	}
	return "???"
}

// Set is a generic unique set of comparable types.
type Set[T comparable] map[T]bool

// Select returns a new List of items from the Set that satisfy the given condition.
func (s Set[T]) Select(f func(T) bool) (selected List[T]) {
	selected = make([]T, 0)
	for item := range s {
		if f(item) {
			selected = append(selected, item)
		}
	}
	return
}

// ToSlice converts the Set to a List, then returns the slice value of it.
func (s Set[T]) ToSlice() []T { return s.ToList().ToSlice() }

// ToMap returns the underlying map.
func (s Set[T]) ToMap() map[T]bool { return s }

// ToSlice converts the Set to a List.
func (s Set[T]) ToList() List[T] {
	if len(s) == 0 {
		return []T{}
	}

	slice := make([]T, len(s))
	i := 0
	for item := range s {
		slice[i] = item
		i++
	}
	return slice
}

// ToSet returns the Set as a Set. This is not useful and only exists to satisfy
// Collection[T].
func (s Set[T]) ToSet() Set[T] { return s }

// GetRandom gets a random T from the Set. This method is not efficient since
// the size of the Set and the index of the chosen random item both affect the
// performance. It is, however, more efficient than converting to a List.
func (s Set[T]) GetRandom() (chosen T) {
	if len(s) == 0 {
		return
	} else {
		n := rand.Intn(len(s))
		i := 0
		for item := range s {
			if i == n {
				chosen = item
				break
			}
			i++
		}
	}
	return
}

// ToShuffledList converts the Set to a List and shuffles it.
func (s Set[T]) ToShuffledList() List[T] { return s.ToList().ToShuffledList() }

// ToSortedList converts the Set to a List and sorts it according to the sort
// function given.
func (s Set[T]) ToSortedList(f func(T, T) int) List[T] { return s.ToList().ToSortedList(f) }

// Len return the length of the underlying Set map.
func (s Set[T]) Len() int { return len(s) }

func (s Set[T]) String() string {
	return s.ToList().String()
}
