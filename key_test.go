package tinykms

import (
	"encoding/json"
	"testing"
)

func TestKey(t *testing.T) {
	rawJSON := []byte(`{ "id": "1", "raw": "1234abcd", "priority": 10 }`)

	var key Key
	err := json.Unmarshal(rawJSON, &key)

	if err != nil {
		t.Errorf("Expected no error. Received: %s", err)
	}

	if key.ID != "1" {
		t.Errorf("Expected ID to be %s, got %s", "1", key.ID)
	}

	if key.Raw != "1234abcd" {
		t.Errorf("Expected Raw to be %s, got %s", "1234abcd", key.Raw)
	}

	if key.Priority != KeyPriority(10) {
		t.Errorf("Expected Priority to be %d, got %d", 10, key.Priority)
	}
}

func TestKeyString(t *testing.T) {
	rawJSON := []byte(`{ "id": "1", "raw": "1234abcd", "priority": 10 }`)

	var key Key
	err := json.Unmarshal(rawJSON, &key)

	if err != nil {
		t.Errorf("Expected no error. Received: %s", err)
	}

	if key.String() != "1234abcd" {
		t.Errorf("Expected 1234abcd, received: %s", key.String())
	}
}
