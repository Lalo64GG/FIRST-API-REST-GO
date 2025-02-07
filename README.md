# My First REST API in Go

## Introduction

This project is a basic implementation of a REST API using Go and the Gin web framework. It serves as a foundation for understanding how to create and manage RESTful services in Go.

## Project Structure

- **`main.go`**: The main entry point of the application, where the server is initialized and routes are defined.
- **`models/album.go`**: Defines the structure of the album data used in the API.
- **`data/data.go`**: Contains sample data of albums.
- **`handlers/handlers.go`**: Contains the handlers for the API endpoints.

## Features

- **Endpoint `/`**: Returns a welcome message to confirm the server is running.
  - **Method**: `GET`
  - **Response**: `Hello world, this is my first API REST in GO`

- **Endpoint `/albums`**: Returns a list of predefined albums in JSON format.
  - **Method**: `GET`
  - **Response**: A JSON array of album objects.

- **Endpoint `/album/:id`**: Returns a specific album based on the provided ID.
  - **Method**: `GET`
  - **Path Parameter**: `id` (ID of the album)
  - **Response**: A JSON object representing the album with the given ID, or a 404 error if not found.

- **Endpoint `/album`**: Adds a new album to the list.
  - **Method**: `POST`
  - **Request Body**: A JSON object representing the new album (e.g., `{ "id": "4", "title": "New Album", "artist": "New Artist", "year": 2024 }`)
  - **Response**: The newly created album in JSON format with a 201 status code.

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
   git clone <https://github.com/Lalo64GG/FIRST-API-REST-GO.git>

4. **Navigate to the Project Directory**:
   ```bash
   cd <FIRST-API-REST-GO>

5. **Run the Application**:
   ```bash
   go run main.go


## Testing the Endpoints

You can test the API using tools like curl or Postman:

1. **Get All Albums**: 
   ```bash	
   curl http://localhost:8080/albums

2. **Get Album by ID**: 
   ```bash
   curl http://localhost:8080/album/1

3. **Post New Album**: 
   ```bash
   curl -X POST http://localhost:8080/album -H "Content-Type: application/json" -d '{"id": "4", "title": "New Album", "artist": "New Artist", "year": 2024}'