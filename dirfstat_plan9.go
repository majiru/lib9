package lib9

import (
	"errors"
	"os"
	"syscall"
)

func Dirfstat(f *os.File) (*syscall.Dir, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if d, ok := fi.Sys().(*syscall.Dir); ok {
		return d, nil
	}

	//This should never happen
	return nil, errors.New("Cast to *syscall.Dir failed")
}
