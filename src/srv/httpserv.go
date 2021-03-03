package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/plcount", plusCounter)
	http.HandleFunc("/debug", debug)
	http.HandleFunc("/gif", gifHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL PATH = %s", r.URL.Path)
}

// Увеличивает счетчик
func plusCounter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Countner++.")
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d", count)
	mu.Unlock()
}

func debug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s, URL: %s, Protp: %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Remote-addr: %s\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func gifHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		fmt.Printf("ParseForm err = %v\n", err)
	}

	p := parametrs{64, 8, 0.001, 5.0, 100.0}

	for k, _ := range r.Form {
		if k == "cycles" {
			n, err := strconv.ParseFloat(r.FormValue(k), 64)
			if err != nil {
				fmt.Printf("ParseFloat err = %v", err)
			}
			p.cycles = n
			fmt.Printf("Cycle = %f", n)
		}
	}

	Lissajous(w, p)
}
