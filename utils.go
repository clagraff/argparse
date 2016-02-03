package parg

import (
	"bytes"
	"strings"

	"github.com/nsf/termbox-go"
)

func getScreenWidth() int {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ := termbox.Size()
	termbox.Close()

	return w
}

func join(delimiter string, args ...string) string {
	var join bytes.Buffer
	num := len(args)

	if num == 0 {
		return ""
	}

	for index, val := range args {
		join.WriteString(val)
		if index < num-1 {
			join.WriteString(delimiter)
		}
	}

	return join.String()
}

func spacer(length int) string {
	count := 0
	char := " "
	var buff bytes.Buffer

	for count < length {
		buff.WriteString(char)
		count++
	}

	return buff.String()
}

func wordWrap(text string, max int) []string {
	var lines []string
	var line []string

	if len(text) <= max {
		return []string{text}
	}

	split := strings.Split(text, " ")
	length := 0

	if len(split) <= 1 {
		return split
	}

	for _, word := range split {
		if len(word)+length+len(line) > max {
			lines = append(lines, join(" ", line...))
			line = []string{word}
			length = len(word)
		} else {
			length = length + len(word)
			line = append(line, word)
		}
	}
	lines = append(lines, join(" ", line...))

	return lines
}
