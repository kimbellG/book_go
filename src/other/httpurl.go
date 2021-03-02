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
	for _, url := range os.Args[1:] {
		startURL := time.Now()
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			fmt.Printf("url: %s\n", url)
		}

		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status code: %d\n", req.StatusCode)
		nbytes, err := io.Copy(io.Discard, req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s. %v", url, err)
			os.Exit(1)
		}
		fmt.Printf("Read %s: %s. Time: %.2fs. Count bytes: %d\n", url, req.Status, time.Since(startURL).Seconds(), nbytes)
	}
	fmt.Printf("Time elapsed: %.2fs", time.Since(start).Seconds())
}
