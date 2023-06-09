package crypto

import (
	"testing"
)

func TestHmacMD5(t *testing.T) {

	k := "123456"
	r := HmacMD5("bluefrog", k)
	e := "20238ad293024e2ea2f505db927cd52e"
	if e != r {
		t.Errorf("e = %v , r = %v ,k = %v", e, r, k)
	}

	r = HmacMD5("admin", k)
	e = "216ae99394e71284ff068b61e422a044"
	if e != r {
		t.Errorf("e = %v , r = %v ,k = %v", e, r, k)
	}
}
