package common

import (
	"slices"
	"testing"
)

func TestListBasics(t *testing.T) {
	input := []rune{'a', 'b', 'c', 'd', 'd'}
	runeList := List[rune](input)

	if toSlice := runeList.ToSlice(); !slices.Equal(input, toSlice) {
		t.Logf("Expected %v to equal %v\n", input, toSlice)
		t.Fail()
	} else {
		if backToRuneSlice := []rune(runeList); !slices.Equal(backToRuneSlice, toSlice) {
			t.Logf("Expected %v to equal %v\n", backToRuneSlice, toSlice)
			t.Fail()
		}
	}

	if runeList.Len() != len(input) {
		t.Log("Expected length of", len(input), "to equal", runeList.Len())
		t.Fail()
	}

	if randomRune := runeList.GetRandom(); !slices.Contains(input, randomRune) {
		t.Log("Expected", randomRune, "to be contained in List")
		t.Fail()
	}

	if toSet := runeList.ToSet(); len(input) == toSet.Len() {
		t.Log("Expected ToSet() to eliminate duplicate!")
		t.Fail()
	}
}

func TestSetBasics(t *testing.T) {
	input := []rune{'a', 'b', 'c', 'd', 'd', 'e', 'e', 'a'}

	if collection := CollectionFrom[rune](input); len(input) != collection.Len() {
		t.Log("Expected length of", len(input), "to equal", collection.Len())
		t.Fail()
	} else {
		if set := collection.ToSet(); !(set.Len() < collection.Len()) {
			t.Log("Expected length of", set.Len(), "to be less than", collection.Len())
			t.Fail()
		}
	}

	if set := List[rune](input).ToSet(); !(set.Len() < len(input)) {
		t.Log("Expected length of", set.Len(), "to be less than", len(input))
		t.Fail()
	} else {
		if random := set.GetRandom(); !slices.Contains(input, random) {
			t.Log(random, "is not in set!")
			t.Fail()
		}
		if list := set.ToList(); list.Len() != set.Len() {
			t.Log("Expected", list.Len(), "to equal", set.Len())
			t.Fail()
		}
	}
}

func TestCollectionSort(t *testing.T) {
	input := []rune{'z', 'g', 'c', 'a', 'b', 'z'}
	sortedListSolution := []rune{'a', 'b', 'c', 'g', 'z', 'z'}
	sortedSetSolution := []rune{'a', 'b', 'c', 'g', 'z'}

	list := CollectionFrom[rune](input)

	runeCompare := func(a, b rune) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	sortedList := list.ToSortedList(runeCompare)
	sortedSet := list.ToSet().ToSortedList(runeCompare)

	if !slices.Equal(sortedListSolution, sortedList) {
		t.Logf("Expected %v to equal %v\n", sortedListSolution, sortedList)
		t.Fail()
	}

	if !slices.Equal(sortedSetSolution, sortedSet) {
		t.Logf("Expected %v to equal %v\n", sortedSetSolution, sortedSet)
		t.Fail()
	}
}

func TestShuffledLists(t *testing.T) {
	input := []rune{'a', 'b', 'c', 'd', 'e', 'e', 'f', 'f'}

	collection := CollectionFrom[rune](input)
	set := collection.ToSet()

	if shuffledList := collection.ToShuffledList(); !(len(input) == shuffledList.Len()) {
		t.Log("Expected length", shuffledList.Len(), "to equal", len(input))
	} else {
		t.Log("The following list should be shuffled. This cannot be programmatically proven, so please verify visually.")
		t.Logf("%c\n", shuffledList)
	}

	if shuffledSet := set.ToShuffledList(); !(shuffledSet.Len() < len(input)) {
		t.Log("Expected length", shuffledSet.Len(), "to be less than", len(input))
		t.Fail()
	} else {
		t.Log("The following list should be shuffled. This cannot be programmatically proven, so please verify visually.")
		t.Logf("%c\n", shuffledSet)
	}
}
