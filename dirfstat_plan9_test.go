package lib9

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDirfstat(t *testing.T) {
	f, err := ioutil.TempFile("", "dirfstat-test")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(f.Name())

	d, err := Dirfstat(f)
	if err != nil {
		t.Fatal(err)
	}

	if filepath.Base(f.Name()) != d.Name {
		t.Fatal("dir does not contains proper name")
	}
}
