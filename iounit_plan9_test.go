package lib9

import (
	"io/ioutil"
	"os"
	"testing"
)

const default9pSize = 8192

func TestIounit(t *testing.T) {
	f, err := ioutil.TempFile("", "iounit-test")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(f.Name())

	n, err := Iounit(f.Fd())
	if err != nil {
		t.Fatal(err)
	}
	if n != default9pSize {
		t.Fatal("iounit return is not expected default size, got:", n)
	}
}
