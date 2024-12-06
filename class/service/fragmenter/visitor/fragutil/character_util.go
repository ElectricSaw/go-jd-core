package fragutil

import (
	"fmt"
)

func EscapeChar(c int) string {
	switch c {
	case '\\':
		return "\\\\"
	case '\b':
		return "\\b"
	case '\f':
		return "\\f"
	case '\n':
		return "\\n"
	case '\r':
		return "\\r"
	case '\t':
		return "\\t"
	case '\'':
		return "\\'"
	case 173: // SOFT HYPHEN
		return unicode(c)
	default:
		if c < ' ' {
			return fmt.Sprintf("\\0%c%c", rune('0'+(c>>3)), rune('0'+(c&7)))
		}
		if c < 127 {
			return string(rune(c))
		}
		if c < 161 {
			return unicode(c)
		}
		return string(rune(c))
	}
}

func unicode(c int) string {
	buffer := make([]rune, 6)
	buffer[0] = '\\'
	buffer[1] = 'u'

	h := c >> 12
	buffer[2] = convertRune(h)
	h = (c >> 8) & 15
	buffer[3] = convertRune(h)
	h = (c >> 4) & 15
	buffer[4] = convertRune(h)
	h = (c) & 15
	buffer[5] = convertRune(h)

	return string(buffer)
}

func convertRune(h int) rune {
	if h <= 9 {
		return rune(h + '0')
	}
	return rune(h + ('A' - 10))
}
