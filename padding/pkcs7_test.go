package padding

import (
	"bytes"
	"testing"
)

// 需要补零 1个
func TestPKCS7PaddingPack(t *testing.T) {
	b := []byte("1234567")
	pack, _ := PKCS7PaddingPack(b, 8)
	e := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte(1)}
	//fmt.Printf("%v = %v", pack, e)
	if !bytes.Equal(e, pack) {
		t.Errorf("Must PKCS7PaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉 1个 填补值
func TestPKCS7PaddingUnPack(t *testing.T) {
	b := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte(1)}
	unpack := PKCS7PaddingUnPack(b)
	e := "1234567"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must PKCS7PaddingUnPack(%v) = %v", b, e)
	}
}

// 需要补零 2个
func TestPKCS7PaddingPack1(t *testing.T) {
	b := []byte("123456")
	pack, _ := PKCS7PaddingPack(b, 8)
	e := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte(2), byte(2)}
	//fmt.Printf("e = %v,pack = %v", e, pack)
	if !bytes.Equal(e, pack) {
		t.Errorf("Must PKCS7PaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉补零 2个
func TestPKCS7PaddingUnPack1(t *testing.T) {
	b := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte(2), byte(2)}
	unpack := PKCS7PaddingUnPack(b)
	e := "123456"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must PKCS7PaddingUnPack(%v) = %v", b, e)
	}
}

// 需要填补 1组
func TestPKCS7PaddingPack2(t *testing.T) {
	b := []byte("12345678")
	pack, _ := PKCS7PaddingPack(b, 8)
	e := []byte{
		byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte('8'),
		byte(8), byte(8), byte(8), byte(8), byte(8), byte(8), byte(8), byte(8),
	}
	//fmt.Printf("e = %v,pack = %v", e, pack)
	if !bytes.Equal(e, pack) {
		t.Errorf("Must PKCS7PaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉 1组 填补
func TestPKCS7PaddingUnPack2(t *testing.T) {
	b := []byte{
		byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte('8'),
		byte(8), byte(8), byte(8), byte(8), byte(8), byte(8), byte(8), byte(8),
	}
	unpack := PKCS7PaddingUnPack(b)
	e := "12345678"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must PKCS7PaddingUnPack(%v) = %v", b, e)
	}
}

// 需要填补 7个
func TestPKCS7PaddingPack3(t *testing.T) {
	b := []byte("1")
	pack, _ := PKCS7PaddingPack(b, 8)
	e := []byte{byte('1'), byte(7), byte(7), byte(7), byte(7), byte(7), byte(7), byte(7)}
	//fmt.Printf("e = %v,pack = %v", e, pack)
	if !bytes.Equal(e, pack) {
		t.Errorf("Must PKCS7PaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉填补 7个
func TestPKCS7PaddingUnPack3(t *testing.T) {
	b := []byte{byte('1'), byte(7), byte(7), byte(7), byte(7), byte(7), byte(7), byte(7)}
	unpack := PKCS7PaddingUnPack(b)
	e := "1"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must PKCS7PaddingUnPack(%v) = %v", b, e)
	}
}
