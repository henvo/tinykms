package tinykms

import (
	"encoding/json"
	"testing"
)

var set KeySet
var rawJSON string = `
		{
			"keys": [
				{ "id": "2", "raw": "1234abce", "priority": 20 },
				{ "id": "3", "raw": "1234abcf", "priority": 50 },
				{ "id": "4", "raw": "1234abc0", "priority": 0 },
				{ "id": "1", "raw": "1234abcd", "priority": 10 }
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

	set, err = FromString("blob")

	if err == nil {
		t.Errorf("Expected error, received nil")
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

	var emptySet KeySet

	_, err = emptySet.GetKeyByID("2")

	if err == nil {
		t.Errorf("Expected error, received nil")
	}

	if err != EmptyKeySetError {
		t.Errorf("Expected EmptyKeySetError, received: %s", err)
	}

	_, err = set.GetKeyByID("non-existing")

	if err == nil {
		t.Errorf("Expected error, received nil")
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

	var emptySet KeySet

	_, err = emptySet.GetKey()

	if err == nil {
		t.Errorf("Expected error, received nil")
	}

	if err != EmptyKeySetError {
		t.Errorf("Expected EmptyKeySetError, received: %s", err)
	}
}
