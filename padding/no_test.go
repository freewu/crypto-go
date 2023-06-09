package padding

import (
	"bytes"
	"testing"
)

func TestNoPaddingPack(t *testing.T) {
	b := []byte("12345678")
	pack, _ := NoPaddingPack(b, 8)
	if !bytes.Equal(b, pack) {
		t.Errorf("Must NoPaddingPack(%v,8) = %v", b, pack)
	}
}

// 抛出异常,不满足分组要求 更长
func TestNoPaddingPackException(t *testing.T) {
	b := []byte("123456789")
	_, err := NoPaddingPack(b, 8)
	
	if err == nil {
		t.Errorf("len(%v) %% 8 = %v,不满足分组,未抛出异常", b, len(b)/8)
	}
}

func TestNoPaddingUnPack(t *testing.T) {
	b := []byte("12345678")
	unpack := NoPaddingUnPack(b)
	if !bytes.Equal(b, unpack) {
		t.Errorf("Must NoPaddingUnPack(%v) = %v", b, unpack)
	}
}
