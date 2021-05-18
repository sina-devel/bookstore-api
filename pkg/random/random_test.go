package random

import "testing"

func TestStringWithCharset(t *testing.T) {
	randomString := StringWithCharset(4, "abcdef")
	if len(randomString) != 4 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	randomString := String(4)
	if len(randomString) != 4 {
		t.Fail()
	}
}
