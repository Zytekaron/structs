package treemap

import (
	"github.com/zytekaron/structs"
	"testing"
)

func TestTreeMap(t *testing.T) {
	tm := NewOrdered[string, string]()
	it := tm.Iterator()

	var _ structs.Map[string, string] = tm // works
	var _ structs.Iterator[string] = it    // works

	if tm.Size() != 0 {
		t.Error("initial size not 0, got", tm.Size())
	}

	tm.Put("hello", "world")
	tm.Put("foo", "bar")

	tm.Put("removed", "woop")
	if tm.Remove("removed") != "woop" {
		t.Error("removed wasn't removed")
	}

	tm.Put("color", "blue")
	tm.Put("model", "T")

	if tm.Size() != 4 {
		t.Error("size not 4, got", tm.Size())
	}

	got := tm.Get("hello")
	if got != "world" {
		t.Errorf("expected 'hello' to map to 'world', got '%s'", got)
	}
	got = tm.Get("foo")
	if got != "bar" {
		t.Errorf("expected 'foo' to map to 'bar', got '%s'", got)
	}
	got = tm.Get("color")
	if got != "blue" {
		t.Errorf("expected 'color' to map to 'blue', got '%s'", got)
	}
	got = tm.Get("model")
	if got != "T" {
		t.Errorf("expected 'model' to map to 'T', got '%s'", got)
	}
}
