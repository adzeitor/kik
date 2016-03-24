package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"
)

var (
	port      int
	bind      string
	maxLength int
)

func init() {
	flag.IntVar(&port, "port", 3000, "web server port(3000 for default)")
	flag.IntVar(&maxLength, "maxlength", 2048, "max length")
	flag.StringVar(&bind, "bind", "127.0.0.1", "binding (127.0.0.1 for default) ")
}

func main() {
	flag.Parse()

	// for compatibility with api.left-pad.io
	// ?str=paddin%27%20oswalt&len=68&ch=@
	http.HandleFunc("/left", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		str := r.FormValue("str")
		len, _ := strconv.Atoi(r.FormValue("len"))

		if len > maxLength {
			len = maxLength
		}

		ch := r.FormValue("ch")
		if ch == "" {
			ch = " "
		}

		leftPad(str, len, ch, w)
	})

	http.HandleFunc("/right", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		str := r.FormValue("str")
		len, _ := strconv.Atoi(r.FormValue("len"))

		if len > maxLength {
			len = maxLength
		}

		ch := r.FormValue("ch")
		if ch == "" {
			ch = " "
		}

		rightPad(str, len, ch, w)
	})

	log.Printf("Served at %s:%d", bind, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", bind, port), nil))
}

func leftPad(str string, len int, ch string, w io.Writer) {
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

func rightPad(str string, len int, ch string, w io.Writer) {
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
