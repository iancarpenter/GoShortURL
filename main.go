package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]
	for _, u := range urls {
		shortened := shortenURL(u)
		fmt.Fprintln(w, "Shortened URL:", shortened)
	}
}

// shortenURL takes a URL string as input, ensures it has a proper scheme (http or https),
// parses it, and returns the host part of the URL. If the URL is invalid, it returns "Invalid URL".
func shortenURL(u string) string {
	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		u = "http://" + u
	}
	parsedURL, err := url.Parse(u)
	if err != nil {
		return "Invalid URL"
	}
	return parsedURL.Host
}
