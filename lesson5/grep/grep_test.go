package grep

import (
	"bufio"
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func compareString(expected, returned string, t *testing.T) {
	t.Helper()

	if strings.Compare(returned, expected) != 0 {
		t.Errorf("Returned: [%v], Expected: [%v]", returned, expected)
	}
}

func compareBool(expected, returned bool, t *testing.T) {
	t.Helper()

	if returned != expected {
		t.Errorf("Returned: [%v], Expected: [%v]", returned, expected)
	}
}
func compareStruct(expected, returned interface{}, t *testing.T) {
	t.Helper()

	if !reflect.DeepEqual(returned, expected) {
		t.Errorf("Returned: [%v], Expected: [%v]", returned, expected)
	}
}
func TestSubstring(t *testing.T) {
	s := newSubstring("color", "text")
	if s.color != "color" && s.str != "text" {
		t.Error(
			"For", s.color,
			"expected", "color",
			"got", s.color,
			"\n",
			"For", s.str,
			"expected", "text",
			"got", s.str,
		)
	}
}

func TestLine(t *testing.T) {
	s := Substring{
		color: "red",
		str:   "text",
	}
	l := Line{
		number:     0,
		substring:  nil,
		IsMatching: false,
	}
	l.addSubstring(s)
	r := l.substring[0]
	compareStruct(s, r, t)
}

func TestNewGrep(t *testing.T) {
	g := NewGrep("normal", "colorMatch")
	compareString("normal", g.colorNormal, t)
	compareString("colorMatch", g.colorMatch, t)
	compareBool(false, g.LineNumber, t)
	compareBool(false, g.Invert, t)
}

func TestMatch(t *testing.T) {
	g := NewGrep("normal", "colorMatch")
	text := []string{"First", "Second line", "Third"}
	r := bufio.NewReader(strings.NewReader(strings.Join(text, "\n")))
	g.Match(r, "line")
	compareBool(true, g.IsMatching(), t)
	compareBool(true, g.lines[1].IsMatching, t)
	compareString("line", g.lines[1].substring[1].str, t)

	g.Match(r, "megaword")
	compareBool(false, g.IsMatching(), t)
	compareBool(true, len(g.lines) == 0, t)

}

func TestPrint(t *testing.T) {
	g := NewGrep("[colorNormal]", "[colorMatch]")
	text := []string{"First", "Second line", "Third"}
	r := bufio.NewReader(strings.NewReader(strings.Join(text, "\n")))
	g.Match(r, "line")

	var returned string
	bWriter := bytes.NewBufferString(returned)
	writer := bufio.NewWriter(bWriter)
	if g.IsMatching() {
		g.Print(writer)
		returned = bWriter.String()
		compareString("[colorNormal]Second [colorMatch]line[colorNormal]\n", returned, t)
	}

}

func TestPrintInvert(t *testing.T) {
	g := NewGrep("[colorNormal]", "[colorMatch]")
	text := []string{"First", "Second line", "Third"}
	r := bufio.NewReader(strings.NewReader(strings.Join(text, "\n")))

	g.Invert = true
	g.Match(r, "line")
	var returned string
	bWriter := bytes.NewBufferString(returned)
	writer := bufio.NewWriter(bWriter)
	g.Print(writer)
	returned = bWriter.String()
	compareString("[colorNormal]First\n", returned, t)
}

func TestWriteLine(t *testing.T) {
	g := NewGrep("[colorNormal]", "[colorMatch]")
	g.LineNumber = true

	var returned string
	bWriter := bytes.NewBufferString(returned)
	writer := bufio.NewWriter(bWriter)

	g.writeLine(writer, Line{
		number: 100,
		substring: []Substring{
			Substring{
				color: "red",
				str:   "example",
			},
		},
		IsMatching: true,
	})
	writer.Flush()
	compareString("100: redexample", bWriter.String(), t)
}
