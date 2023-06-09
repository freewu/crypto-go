package crypto

import (
	"testing"
)

func TestHmacSHA256(t *testing.T) {

	k := "123456"
	r := HmacSHA256("bluefrog", k)
	e := "1609bd1fa22c3501f3f5176dea1842a972d3e4efe2033b777f17d5a7edaa4cfb"
	if e != r {
		t.Errorf("e = %v , r = %v ,k = %v", e, r, k)
	}

	r = HmacSHA256("admin", k)
	e = "69c6fda19329b530a43354615c573bf640de9d59a814d8cf3a9760c2e5c614d8"
	if e != r {
		t.Errorf("e = %v , r = %v ,k = %v", e, r, k)
	}
}
