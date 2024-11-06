package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Sets up an HTTP server that listens on port 8080
// and handles requests to the "/shorten" endpoint.
func main() {
	http.HandleFunc("/shorten", shorten)
	http.ListenAndServe(":8080", nil)
}

// Handles incoming HTTP requests to shorten URLs.
// Extracts the "url" query parameters and processes each one.
func shorten(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]
	for _, u := range urls {
		shortened := shortenURL(u)
		fmt.Fprintf(w, "Shortened URL: %s\n", shortened)
	}
}

// Takes a URL string and returns a shortened version of it.
// If the URL is invalid, meaning it doesn't begin with either http:// or https://
// or can't be parsed it returns "Invalid URL".
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
