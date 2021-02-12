package goytchat

import "encoding/binary"

func vn(val int) []byte {
	if val < 0 {
		panic("val must be >= 0")
	}
	stuff := make([]byte, 8)
	bytesWritten := binary.PutUvarint(stuff, uint64(val))
	return stuff[:bytesWritten]
}

func tp(a, b int, ary []byte) []byte {
	return append(
		vn(b<<3|a),
		ary...,
	)
}

func rs(a int, ary []byte) []byte {
	return append(
		tp(2, a, vn(len(ary))),
		ary...,
	)
}

func nm(a int, ary int) []byte {
	return tp(0, a, vn(ary))
}
