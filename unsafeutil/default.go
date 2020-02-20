package unsafeutil

import (
	"unsafe"
)

// Str2bytes convert string to []byte
func Str2bytes(s string) []byte {
	stringHeader := *(*[2]int)(unsafe.Pointer(&s))
	var sliceHeader [3]int
	sliceHeader[0] = stringHeader[0]
	sliceHeader[1] = stringHeader[1]
	sliceHeader[2] = stringHeader[1]
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

// Bytes2str convert []byte to string
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
