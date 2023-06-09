package crypto

import (
	"testing"
)

func TestMD5(t *testing.T) {
	
	r := MD5("bluefrog")
	e := "56c549a676dbe0c0613121709d6a5ce2"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}

	r = MD5("admin")
	e = "21232f297a57a5a743894a0e4a801fc3"
	if e != r {
		t.Errorf("e = %v , r = %v", e, r)
	}
}
