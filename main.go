package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Define the 'album' struct
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// Sample data of albums
var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2004},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2004},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2004},
}

// Handler for the '/albums' route
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Handler for the root route '/'
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello world, this is my first API REST in GO")
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes and their handlers
	router.GET("/", helloWorld)
	router.GET("/albums", getAlbums)

	// Start the server on port 8080
	router.Run("localhost:8080")
}
