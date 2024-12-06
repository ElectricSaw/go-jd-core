package fragutil

import "fmt"

func EscapeString(s string) string {
	length := len(s)

	for i := 0; i < length; i++ {
		c := s[i]

		if (c == '\\') || (c == '"') || (c < ' ') {
			sb := s[:i]

			for ; i < length; i++ {
				c = s[i]

				switch c {
				case '\\':
					sb += "\\\\"
					break
				case '\b':
					sb += "\\b"
					break
				case '\f':
					sb += "\\f"
					break
				case '\n':
					sb += "\\n"
					break
				case '\r':
					sb += "\\r"
					break
				case '\t':
					sb += "\\t"
					break
				case '"':
					sb += "\\\""
					break
				default:
					if c < ' ' {
						sb += "\\0"
						sb += fmt.Sprintf("%s%c", sb, rune('0'+(c>>3)))
						sb += fmt.Sprintf("%s%c", sb, rune('0'+(c&7)))
					} else {
						sb = fmt.Sprintf("%s%c", sb, c)
					}
				}
			}

			return sb
		}
	}

	return s
}
