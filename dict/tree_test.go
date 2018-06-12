package dict

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	rn := NewTreeDict()
	if rn.count != 0 {
		t.Errorf("Expected initial word count of 0 but got %d", rn.count)
	}
}

func TestSplit(t *testing.T) {
	a := assert{t}
	expectedFirst := byte('b')
	expectedRemainder := "iggles"

	first, remainder := split("biggles")
	a.ByteEquals(expectedFirst, first, "Split First")
	a.StringEquals(expectedRemainder, remainder, "Split Remainder")
}

func TestAddOne(t *testing.T) {
	a := assert{t}
	td := NewTreeDict()

	td.Add("a")
	a.IntEquals(1, td.count, "AddOne count")
}

func TestChildNode(t *testing.T) {
	a := assert{t}
	td := NewTreeDict()

	td.Add("as")

	fn := td.root.children[byte('a')]
	a.ByteEquals(byte('a'), fn.value, "ChildNode.First")
	ln := fn.children[byte('s')]
	a.ByteEquals(byte('s'), ln.value, "ChildNode.Last")
}

func TestIsWord(t *testing.T) {
	a := assert{t}
	td := NewTreeDict()

	td.Add("as")

	a.True(td.IsWord("as"), "IsWord: as")
	a.False(td.IsWord("a"), "IsWord: a")
	a.False(td.IsWord("ask"), "IsWord: ask")
	a.False(td.IsWord("spam"), "IsWord: spam")
}

func TestIsPrefix(t *testing.T) {
	a := assert{t}
	td := NewTreeDict()

	td.Add("ant")

	a.True(td.IsPrefix("an"), "IsPrefix: an")
	a.True(td.IsPrefix("ant"), "IsPrefix: ant")
	a.False(td.IsPrefix("and"), "IsPrefix: and")
	a.False(td.IsPrefix("anti"), "IsPrefix: anti")
}

func TestReconstitute(t *testing.T) {
	a := assert{t}
	td := NewTreeDict()

	td.Add("ant")

	a.StringEquals("ant", td.match("ant").reconstitute(), "Reconstitute ant")
}

func testPath(t *testing.T) {
	//	a := assert{t}
	td := NewTreeDict()

	td.Add("spam")

	expected := []string{"s", "p", "a", "m"}
	actual := td.match("spam").path()

	if !reflect.DeepEqual(expected, actual) {
		t.Error(actual)
	}
}

type assert struct {
	t *testing.T
}

func (a assert) True(expected bool, context string) {
	if !expected {
		a.t.Errorf("%s: expected to be True but was False", context)
	}
}

func (a assert) False(expected bool, context string) {
	if expected {
		a.t.Errorf("%s: expected to be False but was True", context)
	}
}

func (a assert) ByteEquals(expected byte, actual byte, context string) {
	if expected != actual {
		a.t.Errorf("%s: expected %c but got %c", context, expected, actual)
	}
}

func (a assert) IntEquals(expected int, actual int, context string) {
	if expected != actual {
		a.t.Errorf("%s: expected %d but got %d", context, expected, actual)
	}
}

func (a assert) StringEquals(expected string, actual string, context string) {
	if expected != actual {
		a.t.Errorf("%s: expected '%s' but got '%s'", context, expected, actual)
	}
}
