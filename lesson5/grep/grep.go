// Package grep provides ...
package grep

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Grep struct {
	colorNormal string
	colorMatch  string
	LineNumber  bool
	Invert      bool
	lines       []Line
}

type Substring struct {
	color string
	str   string
}

func newSubstring(color string, str string) Substring {
	return Substring{
		color: color,
		str:   str,
	}
}

func NewGrep(colorNormal, colorMatch string) *Grep {
	return &Grep{
		colorNormal: colorNormal,
		colorMatch:  colorMatch,
		LineNumber:  false,
		Invert:      false,
	}

}

type Line struct {
	number     int64
	substring  []Substring
	IsMatching bool
}

func (l *Line) addSubstring(s Substring) {
	l.substring = append(l.substring, s)
}

func (g *Grep) Match(r *bufio.Reader, pattern string) {
	line := 1
	lines := []Line{}
	for {
		rLine, err := r.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		newLine := Line{
			number: int64(line),
		}
		if strings.Contains(rLine, pattern) {
			newLine.IsMatching = true
			for len(rLine) > 0 {
				if index := strings.Index(rLine, pattern); index > 0 {
					lastIndex := index + len(pattern)
					newLine.addSubstring(newSubstring(g.colorNormal, rLine[:index]))
					newLine.addSubstring(newSubstring(g.colorMatch, rLine[index:lastIndex]))
					rLine = rLine[lastIndex:]
				} else {
					if len(newLine.substring) > 0 {
						newLine.addSubstring(newSubstring(g.colorNormal, rLine[:]))
					}
					rLine = rLine[:0]
				}
			}
		} else {
			newLine.IsMatching = false
			newLine.addSubstring(newSubstring(g.colorNormal, rLine[:]))
			rLine = rLine[:0]
		}

		lines = append(lines, newLine)
		line++
	}

	g.lines = lines
}

func (g Grep) IsMatching() bool {
	for _, l := range g.lines {
		if l.IsMatching {
			return true
		}
	}
	return false
}

func (g Grep) Print(writer *bufio.Writer) {
	// writer := bufio.NewWriter(os.Stdout)
	for _, line := range g.lines {
		if !g.Invert {
			if line.IsMatching {
				g.writeLine(writer, line)
			}
		} else {
			if !line.IsMatching {
				g.writeLine(writer, line)
			}
		}
	}

	writer.Flush()
}

func (g Grep) writeLine(writer *bufio.Writer, line Line) {
	if g.LineNumber {
		writer.WriteString(strconv.FormatInt(line.number, 10))
		writer.WriteString(": ")
	}
	for _, s := range line.substring {
		writer.WriteString(s.color)
		writer.WriteString(s.str)
	}
}
