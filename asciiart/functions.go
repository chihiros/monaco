package asciiart

import (
	"bytes"
)

func line(buf *bytes.Buffer, char string, num int) {
	// fmt.Printf("char:%#v\n", char)
	for i := 0; i < num; i++ {
		buf.WriteString(char)
	}
}

func printLine(char string, a []string) string {
	var buf bytes.Buffer
	sum := 0
	for i := range a {
		sum = sum + len([]rune(a[i]))
	}

	if sum < 10 {
		line(&buf, char, 10)
	} else {
		for _, r := range a {
			for b := range []rune(r) {
				// もし全角の漢字、ひらがな、カタカナなら
				if b >= 0x4e00 && b <= 0x9faf {
					line(&buf, char, 2)
				} else {
					line(&buf, char, 1)
				}
			}
		}
	}
	return buf.String()
}
