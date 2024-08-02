package main

import (
    "github.com/gin-gonic/gin"
    "github.com/api-rest-go/src/handlers"
)

func main() {
    // Inicializar el router Gin
    router := gin.Default()

    // Definir rutas y sus handlers
    router.GET("/", func(c *gin.Context) {
        c.String(200, "Hello world, this is my first API REST in GO")
    })
    router.GET("/albums", handlers.GetAlbums)
    router.GET("/album/:id", handlers.GetAlbumByID)
    router.POST("/album", handlers.PostAlbums)

    // Iniciar el servidor en el puerto 8080
    router.Run("localhost:8080")
}
