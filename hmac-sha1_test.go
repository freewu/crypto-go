package crypto

import (
	"testing"
)

func TestHmacSHA1(t *testing.T) {
	
	k := "123456"
	r := HmacSHA1("bluefrog", k)
	e := "fb9a69e0368d707f23d5f93d8d10fedfaa35882e"
	if e != r {
		t.Errorf("e = %v , r = %v ,k = %v", e, r, k)
	}

	r = HmacSHA1("admin", k)
	e = "3c39afa93e0b12c28f1f08b18488ebd4ad2e5858"
	if e != r {
		t.Errorf("e = %v , r = %v ,k = %v", e, r, k)
	}
}
