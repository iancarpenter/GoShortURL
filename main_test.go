package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.com", "example.com"},
		{"https://example.com", "example.com"},
		{"example.com", "Invalid URL"},
		{"ftp://example.com", "Invalid URL"},
		{"", "Invalid URL"},
	}

	for _, test := range tests {
		result := shortenURL(test.input)
		if result != test.expected {
			t.Errorf("shortenURL(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestShorten(t *testing.T) {
	tests := []struct {
		query    string
		expected string
	}{
		{"url=http://example.com", "Shortened URL: example.com\n"},
		{"url=https://example.com", "Shortened URL: example.com\n"},
		{"url=example.com", "Shortened URL: Invalid URL\n"},
		{"url=ftp://example.com", "Shortened URL: Invalid URL\n"},
		{"url=", "Shortened URL: Invalid URL\n"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/shorten?"+test.query, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(shorten)
		handler.ServeHTTP(rr, req)

		if rr.Body.String() != test.expected {
			t.Errorf("shorten() = %q; want %q", rr.Body.String(), test.expected)
		}
	}
}