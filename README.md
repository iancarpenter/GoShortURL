# GoShortURL

GoShortURL is a simple URL shortening service written in Go. It provides an HTTP endpoint to shorten URLs.

## Getting Started

### Prerequisites

- Go 1.23.1 or later

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/GoShortURL.git
    cd GoShortURL
    ```

2. Build the project:
    ```sh
    go build -o goshorturl
    ```

### Environment Variables

The application relies on the following Environment Variables 

1. PG_HOST
2. PG_PORT
3. PG_USER
4. PG_PASSWORD
5. PG_DB_NAME
6. PG_DB_SSLMODE

### Running the Application

To run the application, execute the following command:
```sh
./goshorturl

```
