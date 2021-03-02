package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	//	var filename string
	//	if strings.HasPrefix(url, "http://") {
	//		filename = url[7:] + ".html"
	//	} else {
	//		filename = url + ".html"
	//	}

	//	file, err := os.Create(filename)
	//	if err != nil {
	//		ch <- fmt.Sprintf("Unable to create file: %s. %v", filename, err)
	//		return
	//	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//	defer file.Close()
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Not copy in file!!. filename: %s. %s", url, err)
		return
	}
	err = resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading: %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Read %s: %s. time: %.2fs. bytes: %d", url, resp.Status, secs, nbytes)
}
