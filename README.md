# My First REST API in Go

## Introduction

This project is a basic implementation of a REST API using Go and the Gin web framework. It serves as a foundation for understanding how to create and manage RESTful services in Go.

## Project Structure

- **`main.go`**: The main entry point of the application, where the server is initialized and routes are defined.

- **`album` struct**: Defines the structure of the album data used in the API.

## Features

- **Endpoint `/`**: Returns a welcome message to confirm the server is running.

- **Endpoint `/albums`**: Returns a list of predefined albums in JSON format.

## Technologies Used

- **Go**: The programming language used to build the API. Go is known for its performance and simplicity, making it an excellent choice for building efficient web services.

- **Gin**: A high-performance web framework for Go. Gin is used to handle HTTP requests, define routes, and manage responses.

## How to Run

1. **Install Go**: Ensure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

2. **Install Gin**: Use the Go module system to install Gin by running:
   ```bash
   go get -u github.com/gin-gonic/gin

3. **Clone the Repository**: Clone the repository to your local machine
   ```bash
   git clone <repository-url>

4. **Navigate to the Project Directory**:
  ```bash
  cd <project-directory>

5. **Run the Application**:
  ```bash
  go run main.go
