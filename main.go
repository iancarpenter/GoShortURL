package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// Sets up an HTTP server that listens on port 8080
// and handles requests to the "/shorten" endpoint.
func main() {
	http.HandleFunc("/shorten", shorten)
	http.ListenAndServe(":8080", nil)
}

// Handles incoming HTTP requests to shorten URLs.
// Extracts the "url" query parameters and processes each one.
// Saves the original URL and the shortened URL to the database.
func shorten(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]
	for _, u := range urls {
		shortened := shortenURL(u)
		fmt.Fprintf(w, "Shortened URL: %s\n", shortened)

		err := saveURLToDatabase(u, shortened)
		if err != nil {
			fmt.Printf("Unable to save URL to database " + err.Error())
		}
	}

}

// Takes a URL string and returns a shortened version of it.
// If the URL is invalid, meaning it doesn't begin with either http:// or https://
// or can't be parsed it returns "Invalid URL".
// Otherwise, it returns the host of the URL.
func shortenURL(u string) string {
	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		return "Invalid URL"
	}

	parsedURL, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
		return "Unable to parse URL"
	}

	return parsedURL.Host
}

// Calls the insertURLRecord function to save the original URL and the shortened URL to the database.
func saveURLToDatabase(u string, parsedURL string) error {
	db, err := connectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = insertURLRecord(db, u, parsedURL)
	if err != nil {
		return err
	}

	return nil
}

// Inserts a record into the database with the original URL and the shortened URL.
func insertURLRecord(db *sql.DB, u string, parsedURL string) error {
	const insertSQL string = "INSERT INTO shorten.url_shorten(original_url,shorten_url)  values($1, $2)"

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u, parsedURL)
	if err != nil || res == nil {
		return err
	}
	return nil
}

// Connects to the database using credentials from environment variables.
func connectToDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_DB_SSLMODE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
