package goytchat

import (
	"bytes"
	"testing"
)

// acquired from: unquote(pytchat.paramgen.arcparam.getparam("dWNvlyycWzQ", 0, False, "UCTUHzVzRwN_2x13IWQ9QVNg"))
const okieDokieBoomer = "op2w0wSCARpsQ2pnYURRb0xaRmRPZG14NWVXTlhlbEVxSndvWVZVTlVWVWg2Vm5wU2QwNWZNbmd4TTBsWFVUbFJWazVuRWd0a1YwNTJiSGw1WTFkNlVSb1Q2cWpkdVFFTkNndGtWMDUyYkhsNVkxZDZVU0FCKAAwADgAQABIBFICIAByAggEeAA="

func TestVN(t *testing.T) {
	if !bytes.Equal([]byte{byte(128), byte(4)}, vn(512)) {
		t.Fatalf("vn does not seem to be variable big endian")
	}
}

func TestGetArchivedParam(t *testing.T) {
	b64 := GetArchivedParam("dWNvlyycWzQ", 0, false, "UCTUHzVzRwN_2x13IWQ9QVNg")
	if b64 != okieDokieBoomer {
		t.Fatalf("expected %s, got %s", okieDokieBoomer, b64)
	}
}
