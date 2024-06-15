# URL Shortener Service

A simple URL shortener service written in Go. This service allows users to submit any URL to be shortened and receive a valid shortened URL that will forward the user's request to the original URL.

## Features

- Shorten any valid URL
- Redirect to the original URL using the shortened URL
- JSON responses for errors and successful operations

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or higher)
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerized testing)
- [Git](https://git-scm.com/)

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/arshadsiddique/url-shortner.git
cd url-shortner
```

#### Running Locally
	
1. Install Dependencies:
   ```sh
   go mod download
   ```
2. Run the Application:
   ```sh
   go run main.go
   ```
3. Test the Application:
   ```sh
   go test -v ./...
   ```

#### Using Docker
1. Build the Docker Image:
   ```sh
   docker build -t urlshortener:latest .
   ```
   
2. Run the Docker Container:
   ```sh
   docker run -p 8080:8080 urlshortener:latest
   ```


#### API Endpoints
* Shorten URL:
  ```sh 
  PUT /shorten
  Content-Type: application/json

  Request Body:
  {
     "destination": "https://www.google.com"
  }
    
  Response Body:
  {
     "shortened_url": "http://localhost:8080/abc123"
  }
  ```

* Redirect to Original URL:
  ```sh
  GET /{shortcode}

  Example:
  GET /abc123
  ```

#### Testing the Endpoints
* Shorten a URL:
  ```sh
  curl -X PUT http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"destination": "https://www.google.com"}'
  ```

* Redirect to the Original URL:
  Assuming the shortened URL returned is http://localhost:8080/abc123, you can use:
  ```sh
  curl -L http://localhost:8080/abc123
  ```
