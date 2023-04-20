package shellwords

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
)

var (
	reSplit     = regexp.MustCompile(`(?m)^\s*(?:(?P<word>[^\s\\'"]+)|'(?P<sq>[^']*)'|"(?P<dq>(?:[^"\\]|\\.)*)"|(?P<esc>\\.?)|(?P<garbage>\S))(?P<sep>\s+|\z)?`)
	reEscapeDq  = regexp.MustCompile(`\\([$"\\\n` + "`" + `])`)
	reEscapeEsc = regexp.MustCompile(`\\(.)`)
)

// Split a string into an array of tokens in the same way the UNIX Bourne shell does.
func Split(line string) ([]string, error) {
	var words []string
	var field bytes.Buffer
	for line != "" {
		m := reSplit.FindStringSubmatch(line)
		if len(m) != 7 {
			return nil, errors.New("unmatched")
		}
		var (
			word    = m[1]
			sq      = m[2]
			dq      = m[3]
			esc     = m[4]
			garbage = m[5]
			sep     = m[6]
		)
		if garbage != "" {
			return nil, fmt.Errorf("unmatched quote: %s", line)
		}
		switch {
		case word != "":
			field.WriteString(word)
		case sq != "":
			field.WriteString(sq)
		case dq != "":
			dq = reEscapeDq.ReplaceAllString(dq, "$1")
			field.WriteString(dq)
		case esc != "":
			esc = reEscapeEsc.ReplaceAllString(esc, "$1")
			field.WriteString(esc)
		}
		line = line[len(m[0]):]
		if sep != "" || line == "" {
			words = append(words, field.String())
			field.Reset()
		}
	}
	return words, nil
}

// Join builds a command line string from an argument list by joining
// all elements escaped for Bourne shell and separated by a space.
func Join(words []string) string {
	var buf bytes.Buffer
	for i, w := range words {
		if i != 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(Escape(w))
	}
	return buf.String()
}

// Escape escapes a string so that it can be safely used in a Bourne shell command line.
// Note that a resulted string should be used unquoted and is not intended for use in double quotes nor in single quotes.
func Escape(str string) string {
	if str == "" {
		return "''"
	}

	strBytes := []byte(str)
	var buf bytes.Buffer
	for _, b := range strBytes {
		switch b {
		case
			'a', 'b', 'c', 'd', 'e', 'f', 'g',
			'h', 'i', 'j', 'k', 'l', 'm', 'n',
			'o', 'p', 'q', 'r', 's', 't', 'u',
			'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G',
			'H', 'I', 'J', 'K', 'L', 'M', 'N',
			'O', 'P', 'Q', 'R', 'S', 'T', 'U',
			'V', 'W', 'X', 'Y', 'Z',
			'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'_', '-', '.', ',', ':', '/', '@':
			buf.WriteByte(b)
		case '\n':
			buf.WriteString("'\n'")
		default:
			buf.WriteByte('\\')
			buf.WriteByte(b)
		}
	}

	return buf.String()
}
