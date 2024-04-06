# tinykms
A tiny in-memory key management "system" that helps managing multiple keys
inside a key set (e.g. rotating).

## Foreword
This project does not claim to be production-ready. It should not be used in
critical applications. If you have any feedback and improvement suggestions
I am happy to hear your feedback or receive a PR. Use at your own risk.

## Usage
``` go
import (
  "fmt"

  "github.com/henvo/tinykms"
)

rawJson := `{
  "keys": [
    { "id": "1", "raw": "...", "priority": 40 },
    { "id": "2", "raw": "...", "priority": 20 },
    { "id": "3", "raw": "...", "priority": 10 },
    { "id": "4", "raw": "...", "priority": 0 }
  ]
}`

// Initialize the key set from string.
ks, _ := tinykms.FromString(rawJson)

// Fetch a key by given ID.
key1, _ := ks.GetKeyByID("1")
fmt.Println(key1.ID) // Prints 1

// Fetch the key with highest priority (lowest number).
key2, _ := ks.GetKey()
fmt.Println(key2.ID) // Prints 3
```
