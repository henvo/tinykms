package tinykms

import (
	"encoding/json"
	"testing"
)

var set KeySet
var rawJSON string = `
		{
			"keys": [
				{ "id": "1", "raw": "1234abcd", "priority": 10 },
				{ "id": "2", "raw": "1234abce", "priority": 20 },
				{ "id": "3", "raw": "1234abcf", "priority": 50 },
				{ "id": "4", "raw": "1234abc0", "priority": 0 }
			]
		}`

func TestKeySet(t *testing.T) {
	err := json.Unmarshal([]byte(rawJSON), &set)

	if err != nil {
		t.Errorf("Expected no error. Received: %s", err)
	}

	if len(set.Keys) != 4 {
		t.Errorf("Expected 4 keys, got %d", len(set.Keys))
	}
}

func TestKeySetFromString(t *testing.T) {
	set, err := FromString(rawJSON)

	if err != nil {
		t.Errorf("Expected no error. Received: %s", err)
	}

	if len(set.Keys) != 4 {
		t.Errorf("Expected 4 keys, got %d", len(set.Keys))
	}
}

func TestKeySetGetKeyByID(t *testing.T) {
	value, err := set.GetKeyByID("2")

	if err != nil {
		t.Errorf("Expected no error. Received: %s", err)
	}

	if value.ID != "2" {
		t.Errorf("Expected key with ID %s, received: %s", "2", value.ID)
	}
}

func TestKeySetGetKey(t *testing.T) {
	value, err := set.GetKey()

	if err != nil {
		t.Errorf("Expected no error. Received: %s", err)
	}

	if value.ID != "1" {
		t.Errorf("Expected key with ID %s, received: %s", "1", value.ID)
	}
}
