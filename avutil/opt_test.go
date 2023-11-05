package avutil

import "testing"

func TestDictionary_Set(t *testing.T) {
	dict := &Dictionary{}
	defer dict.Free()

	count := dict.Count()
	if count != 0 {
		t.Errorf("expected empty dictionary, got %d element(s)", count)
	}

	if !dict.Set("test", "test", 0) {
		t.Errorf("expected success")
	}

	count = dict.Count()
	if count != 1 {
		t.Errorf("expected 1 element, got %d element(s)", count)
	}
}
