package genericsafemap

import "testing"

func TestNewGenericSafeMap(t *testing.T) {
	m := New[string, int]()

	if m == nil {
		t.Errorf("Expected value to be not nil")
	}

	val, ok := m.Get("one")
	if ok || val != 0 {
		t.Errorf("Expected value false and 0, but got true and %v", val)
	}
}

func TestGet(t *testing.T) {
	m := New[string, int]()

	m.Put("one", 1)

	val, ok := m.Get("one")
	if !ok || val != 1 {
		t.Errorf("Expected value to be 1, got %d", val)
	}

	_, ok = m.Get("two")
	if ok {
		t.Errorf("Expected value to be false, got %t", ok)
	}
}

func TestPut(t *testing.T) {
	m := New[string, int]()

	m.Put("one", 1)

	val, ok := m.Get("one")
	if !ok || val != 1 {
		t.Errorf("Expected value 1, but got %v", val)
	}

	m.Put("one", 2)

	val, ok = m.Get("one")
	if !ok || val != 2 {
		t.Errorf("Expected value 2, but got %v", val)
	}
}

func TestRemove(t *testing.T) {
	m := New[string, int]()

	m.Put("one", 1)

	val, ok := m.Get("one")
	if !ok || val != 1 {
		t.Errorf("Expected value 1, but got %v", val)
	}

	m.Remove("one")

	_, ok = m.Get("one")
	if ok {
		t.Errorf("Expected value false, but got true")
	}
}
