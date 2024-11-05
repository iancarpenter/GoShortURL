package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/shorten", shorten)
	http.ListenAndServe(":8080", nil)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]
	for _, u := range urls {
		shortened := shortenURL(u)
		fmt.Fprintf(w, "Shortened URL: %s\n", shortened)
	}
}

func shortenURL(u string) string {
	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		return "Invalid URL"
	}
	parsedURL, err := url.Parse(u)
	if err != nil {
		return "Invalid URL"
	}
	return parsedURL.Host
}
