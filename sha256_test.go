package crypto

import (
	"testing"
)

func TestSHA256(t *testing.T) {
	r := SHA256("bluefrog")
	e := "b374dfd9bcd389c4a6796301382fed0a4b6048fe720c8d96f9b5cbfe58182790"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}

	r = SHA256("admin")
	e = "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}
}
