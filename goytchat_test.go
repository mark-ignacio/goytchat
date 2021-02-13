package goytchat

import (
	"bytes"
	"testing"
)

// acquired from pytchat 0.5.3
const (
	// unquote(pytchat.paramgen.arcparam.getparam("dWNvlyycWzQ", 0, False, "UCTUHzVzRwN_2x13IWQ9QVNg"))
	okieDokieBoomerArchived = "op2w0wSCARpsQ2pnYURRb0xaRmRPZG14NWVXTlhlbEVxSndvWVZVTlVWVWg2Vm5wU2QwNWZNbmd4TTBsWFVUbFJWazVuRWd0a1YwNTJiSGw1WTFkNlVSb1Q2cWpkdVFFTkNndGtWMDUyYkhsNVkxZDZVU0FCKAAwADgAQABIBFICIAByAggEeAA="
	// random.seed(42)
	// unquote(pytchat.paramgen.liveparam.getparam("dWNvlyycWzQ", "UCTUHzVzRwN_2x13IWQ9QVNg", 0, False))
	okieDokieBoomerLive = "0ofMyAPFARpsQ2pnYURRb0xaRmRPZG14NWVXTlhlbEVxSndvWVZVTlVWVWg2Vm5wU2QwNWZNbmd4TTBsWFVUbFJWazVuRWd0a1YwNTJiSGw1WTFkNlVSb1Q2cWpkdVFFTkNndGtWMDUyYkhsNVkxZDZVU0FCKPfU3cD55e4CMAA4AEABShsIABAAGAAgADoAQABKAFDx0dDB-eXuAlgDeABQlcTjwfnl7gJY6cWe5PTl7gJoAYIBAggBiAEAmgECCACgAeKKpsH55e4C"
)

func TestVN(t *testing.T) {
	if !bytes.Equal([]byte{byte(128), byte(4)}, vn(512)) {
		t.Fatalf("vn does not seem to be variable big endian")
	}
}

func TestGetArchivedParam(t *testing.T) {
	b64 := GetArchivedParam("dWNvlyycWzQ", "UCTUHzVzRwN_2x13IWQ9QVNg", 0, false)
	if b64 != okieDokieBoomerArchived {
		t.Fatalf("expected %s, got %s", okieDokieBoomerArchived, b64)
	}
}
func TestGetLiveParam(t *testing.T) {
	// rand.Seed(42)
	b64 := GetLiveParam("dWNvlyycWzQ", "UCTUHzVzRwN_2x13IWQ9QVNg", 0, false)
	// match the 5 random parts.
	b64 = b64[:157] + okieDokieBoomerLive[157:165] + b64[165:] // ts2
	b64 = b64[:198] + okieDokieBoomerLive[198:206] + b64[206:] // ts1
	b64 = b64[:216] + okieDokieBoomerLive[216:223] + b64[223:] // ts3
	b64 = b64[:228] + okieDokieBoomerLive[228:236] + b64[236:] // ts4
	b64 = b64[:261] + okieDokieBoomerLive[261:269] + b64[269:] // ts5
	if b64 != okieDokieBoomerLive {
		t.Fatalf("expected %s, got %s", okieDokieBoomerLive, b64)
	}
}
