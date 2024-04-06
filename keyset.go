package tinykms

import (
	"encoding/json"
	"errors"
	"fmt"
)

var EmptyKeySetError = errors.New("KeySet is empty")

// KeySet holds all the keys provided.
type KeySet struct {
	// Keys hold all keys.
	Keys []Key
}

// FromBytes initializes a KeySet from a JSON byte slice.
func FromBytes(rawJSON []byte) (*KeySet, error) {
	var ks KeySet
	err := json.Unmarshal(rawJSON, &ks)

	if err != nil {
		return &ks, fmt.Errorf("Could not init keyset: %w", err)
	}

	return &ks, nil
}

// FromString initializes a KeySey from a JSON string.
func FromString(rawJSON string) (ks *KeySet, err error) {
	return FromBytes([]byte(rawJSON))
}

// GetKey returns a key by given ID.
func (ks *KeySet) GetKeyByID(id string) (*Key, error) {
	if len(ks.Keys) == 0 {
		return &Key{}, EmptyKeySetError
	}

	for _, k := range ks.Keys {
		if k.ID == id {
			return &k, nil
		}
	}

	return &Key{}, fmt.Errorf("No key found for ID: %s", id)
}

// GetKey selects the Key with the highest priority. The priority is
// defined by the smallest number for priority. Zero is no priority at all.
// If multiple keys have the same high priority the first key is picked.
func (ks *KeySet) GetKey() (*Key, error) {
	if len(ks.Keys) == 0 {
		return &Key{}, EmptyKeySetError
	}

	highestPrioKey := &ks.Keys[0]

	for i, k := range ks.Keys {
		if k.Priority == KeyPriority(0) {
			continue
		}

		if k.Priority < highestPrioKey.Priority {
			highestPrioKey = &ks.Keys[i]
		}
	}

	return highestPrioKey, nil
}
