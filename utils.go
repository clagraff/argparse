package parg

import "bytes"

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
