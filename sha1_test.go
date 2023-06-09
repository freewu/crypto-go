package crypto

import (
	"testing"
)

func TestSHA1(t *testing.T) {

	r := SHA1("bluefrog")
	e := "115c46aef40a5232486db602cb1cab293b85691c"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}

	r = SHA1("admin")
	e = "d033e22ae348aeb5660fc2140aec35850c4da997"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}
}
