package persistence

import (
	"testing"
)

func TestSetAndGet(t *testing.T) {
	p := NewPersistence()
	key := "testKey"
	value := "testValue"

	p.Set(key, value)
	got, ok := p.Get(key)
	if !ok {
		t.Errorf("Expected to get value for key %s, but got nothing", key)
	}
	if got != value {
		t.Errorf("Expected value %s, but got %s", value, got)
	}
}

func TestDelete(t *testing.T) {
	p := NewPersistence()
	key := "testKey"
	value := "testValue"

	p.Set(key, value)
	p.Delete(key)
	_, ok := p.Get(key)
	if ok {
		t.Errorf("Expected key %s to be deleted, but it still exists", key)
	}
}

func TestGetAll(t *testing.T) {
	p := NewPersistence()
	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	for k, v := range data {
		p.Set(k, v)
	}

	all := p.GetAll()
	for k, v := range data {
		if all[k] != v {
			t.Errorf("Expected key %s to have value %s, but got %s", k, v, all[k])
		}
	}
}
