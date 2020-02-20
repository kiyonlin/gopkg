package unsafeutil

import (
	"testing"
)

func TestStr2bytes(t *testing.T) {
	str := "foo"
	b := Str2bytes(str)
	if str != string(b) {
		t.Errorf("convert string to []byte failed! want %s got %s", str, string(b))
	}
}

func TestBytes2str(t *testing.T) {
	b := []byte("foo")
	str := Bytes2str(b)
	if str != string(b) {
		t.Errorf("convert []byte to string failed! want %s got %s", string(b), str)
	}
}
