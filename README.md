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
