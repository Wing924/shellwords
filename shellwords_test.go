package shellwords

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	expected []string
	message  string
}

func TestSplitString(t *testing.T) {
	testCases := map[string]testCase{
		`a "b b" a`:      {[]string{"a", "b b", "a"}, "quoted strings"},
		`a "'b' c" d`:    {[]string{"a", "'b' c", "d"}, "escaped double quotes"},
		`a '"b" c' d`:    {[]string{"a", `"b" c`, "d"}, "escaped single quotes"},
		`a b\ c d`:       {[]string{"a", "b c", "d"}, "escaped spaces"},
		`a   b\ c d`:     {[]string{"a", "b c", "d"}, "extra spaces in seplator"},
		`   a b\ c d`:    {[]string{"a", "b c", "d"}, "extra leading spaces"},
		`a b\ c d   `:    {[]string{"a", "b c", "d"}, "extra tailing spaces"},
		"a 'aa\nbb\ncc'": {[]string{"a", "aa\nbb\ncc"}, "multi-line"},
	}
	errorCases := []string{
		`a "b c d e`,
		`a 'b c d e`,
		`"a "'b' c" d`,
	}

	for input, res := range testCases {
		actual, err := Split(input)
		assert.NoError(t, err)
		assert.Equal(t, res.expected, actual, res.message)
	}
	for _, input := range errorCases {
		_, err := Split(input)
		assert.Error(t, err)
	}
}

func TestEscape(t *testing.T) {
	testCases := []string{
		``,
		`abc`,
		`a b c`,
		`a  b `,
		`a\nb`,
		"a\nb",
		"a\n\nb",
		`a $HOME`,
		`sh -c 'pwd'`,
		`a"b'`,
	}

	for _, expected := range testCases {
		escaped := Escape(expected)
		actual, err := exec.Command("sh", "-c", "printf %s "+escaped).Output()
		assert.NoError(t, err)
		assert.Equal(t, expected, string(actual), "input: [%s], escaped: [%s], actual: [%s]", expected, escaped, actual)
	}
}

func TestJoin(t *testing.T) {
	testCases := map[string][]string{
		"":                {},
		"a b c":           {"a", "b", "c"},
		`a\ b c`:          {"a b", "c"},
		`sh -c echo\ foo`: {"sh", "-c", "echo foo"},
	}
	for expected, input := range testCases {
		actual := Join(input)
		assert.Equal(t, expected, actual, "input: %#v, expected: (%s), actual: (%s)", input, expected, actual)
	}
}
