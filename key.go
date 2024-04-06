package tinykms

import "time"

// KeyPriority is an indicator how a key should be handled by the KMS.
// 0 = no priority
// 1 = highest priority
// 255 = lowest priority
type KeyPriority uint8

// Key is the representation for a single key with its metadata.
type Key struct {
	// ID is unique identifier for the given (raw) key.
	ID string `json:"id"`

	// Raw is the key itself.
	Raw string `json:"raw"`

	// Priority allows easier and more graceful key rotation.
	Priority KeyPriority `json:"priority"`

	// CreatedAt represents the date the key was generated.
	CreatedAt time.Time `json:"created_at"`
}
