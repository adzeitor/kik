package main

import (
	"log"
	"net/http"
	"strconv"
	"os"
	"encoding/json"
)

var (
	port      int
	bind      string
	maxLength int
)

type Answer struct {
	Str string `json:"str"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	bind := os.Getenv("BIND")

	maxLength,_ := strconv.Atoi(os.Getenv("MAXLENGTH"))
	if maxLength <= 0 {
		// free version limit
		maxLength = 2048
	}

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

		FLeftPad(w, str, len, ch)
	})

	http.HandleFunc("/left.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		str := r.FormValue("str")
		len, _ := strconv.Atoi(r.FormValue("len"))

		if len > maxLength {
			len = maxLength
		}

		ch := r.FormValue("ch")
		if ch == "" {
			ch = " "
		}

		padded := LeftPad(str, len, ch)

		res, _ := json.Marshal(Answer{padded})
		w.Write(res)
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

		FRightPad(w, str, len, ch)
	})

	http.HandleFunc("/right.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		str := r.FormValue("str")
		len, _ := strconv.Atoi(r.FormValue("len"))

		if len > maxLength {
			len = maxLength
		}

		ch := r.FormValue("ch")
		if ch == "" {
			ch = " "
		}

		padded := RightPad(str, len, ch)

		res, _ := json.Marshal(Answer{padded})
		w.Write(res)
	})

	log.Printf("Served at %s:%s", bind, port)
	log.Fatal(http.ListenAndServe(bind + ":" + port, nil))
}
