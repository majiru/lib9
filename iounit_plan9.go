package lib9

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fdctlArgs = 10

var IounitZeroRead = errors.New("0 read of fd ctl")

var IounitBadRead = errors.New("Bad read of fd ctl")

func Iounit(fd uintptr) (uint64, error) {
	f, err := os.Open(fmt.Sprintf("#d/%dctl", fd))
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 127)
	n, err := f.Read(buf)
	f.Close()
	if err != nil {
		return 0, err
	}

	if n == 0 {
		return 0, IounitZeroRead
	}

	args := Tokenize(string(buf))

	if len(args) < 8 {
		return 0, IounitBadRead
	}

	return strconv.ParseUint(args[7], 10, 64)
}
