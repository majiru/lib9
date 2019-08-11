package lib9

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	var expected []string

	arg := `Hello World`
	expected = []string{`Hello`, `World`}
	fields := Tokenize(arg)
	if fields[0] != expected[0] || fields[1] != expected[1] {
		t.Fatalf("Expected %v got %v", expected, fields)
	}

	arg = `test 'test 123'`
	expected = []string{`test`, `test 123`}
	fields = Tokenize(arg)
	if fields[0] != expected[0] || fields[1] != expected[1] {
		t.Fatalf("Expected %v got %v", expected, fields)
	}

	arg = `test 'test ''123'`
	expected = []string{`test`, `test '123`}
	fields = Tokenize(arg)
	if fields[0] != expected[0] || fields[1] != expected[1] {
		t.Fatalf("Expected %v got %v", expected, fields)
	}
}
