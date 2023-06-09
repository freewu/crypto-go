package padding

import (
	"bytes"
	"testing"
)

// 需要补零 1个
func TestZeroPaddingPack(t *testing.T) {
	b := []byte("1234567")
	pack, _ := ZeroPaddingPack(b, 8)
	e := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte(0)}
	if !bytes.Equal(e, pack) {
		t.Errorf("Must ZeroPaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉补零 1个
func TestZeroPaddingUnPack(t *testing.T) {
	b := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte(0)}
	unpack := ZeroPaddingUnPack(b)
	e := "1234567"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must ZeroPaddingUnPack(%v) = %v", b, e)
	}
}

// 需要补零 2个
func TestZeroPaddingPack1(t *testing.T) {
	b := []byte("123456")
	pack, _ := ZeroPaddingPack(b, 8)
	e := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte(0), byte(0)}
	//fmt.Printf("e = %v,pack = %v", e, pack)
	if !bytes.Equal(e, pack) {
		t.Errorf("Must ZeroPaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉补零 2个
func TestZeroPaddingUnPack1(t *testing.T) {
	b := []byte{byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte(0), byte(0)}
	unpack := ZeroPaddingUnPack(b)
	e := "123456"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must ZeroPaddingUnPack(%v) = %v", b, e)
	}
}

// 需要补零 一组
func TestZeroPaddingPack2(t *testing.T) {
	b := []byte("12345678")
	pack, _ := ZeroPaddingPack(b, 8)
	e := []byte{
		byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte('8'),
		byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0),
	}
	//fmt.Printf("e = %v,pack = %v", e, pack)
	if !bytes.Equal(e, pack) {
		t.Errorf("Must ZeroPaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉补零 1组
func TestZeroPaddingUnPack2(t *testing.T) {
	b := []byte{
		byte('1'), byte('2'), byte('3'), byte('4'), byte('5'), byte('6'), byte('7'), byte('8'),
		byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0),
	}
	unpack := ZeroPaddingUnPack(b)
	e := "12345678"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must ZeroPaddingUnPack(%v) = %v", b, e)
	}
}

// 需要补零 7
func TestZeroPaddingPack3(t *testing.T) {
	b := []byte("1")
	pack, _ := ZeroPaddingPack(b, 8)
	e := []byte{byte('1'), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0)}
	//fmt.Printf("e = %v,pack = %v", e, pack)
	if !bytes.Equal(e, pack) {
		t.Errorf("Must ZeroPaddingPack(%v,8) = %v", b, e)
	}
}

// 需要去掉补零 7个
func TestZeroPaddingUnPack3(t *testing.T) {
	b := []byte{byte('1'), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0), byte(0)}
	unpack := ZeroPaddingUnPack(b)
	e := "1"
	//fmt.Printf("%v = %v", string(unpack), e)
	if string(unpack) != e {
		t.Errorf("Must ZeroPaddingUnPack(%v) = %v", b, e)
	}
}
