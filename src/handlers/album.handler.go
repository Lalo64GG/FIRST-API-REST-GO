package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/api-rest-go/src/data"
    "github.com/api-rest-go/src/models"
)

// Handler para la ruta '/albums'
func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, data.Albums)
}

// Handler para la ruta '/album/:id'
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range data.Albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such album"})
}

// Handler para la ruta '/album' (POST)
func PostAlbums(c *gin.Context) {
    var newAlbum models.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    data.Albums = append(data.Albums, newAlbum)

    c.IndentedJSON(http.StatusCreated, newAlbum)
}
