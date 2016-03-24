package main

import (
	"bytes"
	"io"
	"unicode/utf8"
)

func LeftPad(str string, len int, ch string) string {
	var buf bytes.Buffer

	strlen := utf8.RuneCountInString(str)

	for {
		for _, x := range ch {
			if strlen >= len {
				buf.WriteString(str)
				return buf.String()
			}

			buf.WriteRune(x)
			strlen++
		}
	}
}

func RightPad(str string, len int, ch string) string {
	var buf bytes.Buffer

	strlen := utf8.RuneCountInString(str)
	buf.WriteString(str)

	for {
		for _, x := range ch {
			if strlen >= len {
				return buf.String()
			}

			buf.WriteRune(x)
			strlen++
		}
	}
}

func FRightPad(w io.Writer, str string, len int, ch string) {
	strlen := utf8.RuneCountInString(str)

	w.Write([]byte(str))

	buf := make([]byte, 3)
	for {
		for _, x := range ch {
			if strlen >= len {
				return
			}

			utf8.EncodeRune(buf, x)
			w.Write(buf)
			strlen++
		}
	}
}

func FLeftPad(w io.Writer, str string, len int, ch string) {
	strlen := utf8.RuneCountInString(str)

	buf := make([]byte, 3)
	for {
		for _, x := range ch {
			if strlen >= len {
				w.Write([]byte(str))
				return
			}

			utf8.EncodeRune(buf, x)
			w.Write(buf)
			strlen++
		}
	}
}
