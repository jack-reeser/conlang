package morph

import "testing"

func TestMorpheme(t *testing.T) {
	for _, testCase := range []struct {
		Input  [2]Morpheme
		Output Morpheme
	}{
		{[2]Morpheme{NewPrefix("un"), NewStem("do")}, NewStem("undo")},
		{[2]Morpheme{NewStem("couch"), NewStem("pillow")}, NewStem("couchpillow")},
		{[2]Morpheme{NewSuffix("ly"), NewStem("mad")}, NewStem("madly")},
		{[2]Morpheme{NewPrefix("un"), NewPrefix("re")}, NewPrefix("unre")},
		{[2]Morpheme{NewSuffix("ly"), NewPrefix("re")}, NewStem("rely")},
		{[2]Morpheme{NewSuffix("ly"), NewSuffix("ly")}, NewSuffix("lyly")},
	} {
		newMorpheme := testCase.Input[0].Combine(testCase.Input[1])
		t.Log(newMorpheme)
		if newMorpheme.String() != testCase.Output.String() {
			t.Logf("Expected new Morpheme %s to equal %s\n", newMorpheme, testCase.Output)
			t.Fail()
		}
		if newMorpheme.IsFree() != testCase.Output.IsFree() {
			t.Logf("Expected new Morpheme.IsFree() to equal %t; got %t\n", testCase.Output.IsFree(), newMorpheme.IsFree())
			t.Fail()
		}
		if newMorpheme.IsPrefix() != testCase.Output.IsPrefix() {
			t.Logf("Expected new Morpheme.IsPrefix() to equal %t; got %t\n", testCase.Output.IsPrefix(), newMorpheme.IsPrefix())
			t.Fail()
		}
	}
}
