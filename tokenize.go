package lib9

import "strings"

const qsep = " \t\r\n"

//See /sys/src/libc/port/tokenize.c:/^qtoken/
func qtoken(s []rune, sep string) (string, []rune) {
	quoting := false
	var b strings.Builder
	var i int

	for i = 0; i < len(s) && (quoting || !strings.ContainsRune(sep, s[i])); {
		if s[i] != '\'' {
			b.WriteRune(s[i])
			i++
			continue
		}
		if !quoting {
			quoting = true
			i++
			continue
		}
		if i+1 == len(s) {
			break
		}
		if s[i+1] != '\'' {
			quoting = false
			i++
			continue
		}
		i++
		b.WriteRune(s[i])
	}
	return b.String(), s[i:]
}

//See /sys/src/libc/port/tokenize.c:/^tokenize/
func Tokenize(s string) (out []string) {
	var field string
	var next []rune = []rune(s)
	for len(next) != 0 {
		field, next = qtoken([]rune(strings.TrimLeft(string(next), qsep)), qsep)
		if field != "" {
			out = append(out, field)
		}
	}
	return
}
