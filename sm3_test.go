package crypto

import (
	"testing"
)

func TestSM3(t *testing.T) {
	r := SM3("bluefrog")
	e := "197fdc824b4a3b3a4749c33bc716b9fc29237911cf18e503624b3584d77f28c1"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}

	r = SM3("admin")
	e = "dc1fd00e3eeeb940ff46f457bf97d66ba7fcc36e0b20802383de142860e76ae6"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}
}
