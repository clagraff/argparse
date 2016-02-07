package parg

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/nsf/termbox-go"
)

// extractFlags will extract all flags from the slice of arguments provided,
// returning one slice of just the invididual flags, and all other arguments
// in a second slice.
func extractFlags(allArgs ...string) (flags, args []string) {
	count := 0
	max := len(allArgs)

	for count < max {
		a := allArgs[count]

		// If we have flag-escape string, assume the next arg is supposed
		// to be normal text instead of potentially being a flag.
		if a == "--" && len(allArgs) > count+1 {
			args = append(args, allArgs[count+1])
			count = count + 2
			continue
		}

		// Using a flag regex, check if we have a normal param or a flag.
		flagRegex := regexp.MustCompile(`^-{1,2}[a-zA-Z]+$`)
		if !flagRegex.MatchString(a) {
			args = append(args, a)
			count++
			continue
		}

		// Okay, we must have a flag. Which type?
		isShort := true
		if len(a) > 2 && a[:2] == "--" {
			isShort = false
		}

		// If short-flag, grab all letters individual flags.
		if isShort == true {
			for _, c := range a[1:] {
				flags = append(flags, string(c))
			}
		} else {
			flags = append(flags, a[2:])
		}
		count++
	}

	return flags, args
}

func getFlagName(flag string) string {
	if len(flag) == 2 && flag[0] == '-' {
		return string(flag[1])
	} else if len(flag) > 2 && flag[0:2] == "--" {
		return flag[2:]
	}
	return flag
}

func getScreenWidth() int {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ := termbox.Size()
	termbox.Close()

	return w
}

func isFlagFormat(text string) bool {
	if len(text) == 0 {
		return false
	}

	if text[0] != '-' {
		return false
	}

	text = text[1:]
	if text[0] == '-' {
		text = text[1:]
	}

	return !strings.Contains(text, "-")
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
