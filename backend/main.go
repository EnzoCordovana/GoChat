/*
 * Fait avec la doc:
 * https://go.dev/doc/tutorial/web-service-gin
 */

package main

// Import des blibliothèques
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// La tructure d'un album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

// Notre slice album
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Fonction pour obtenir notre slice
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	// Setup du serveur
	router := gin.Default()

	// Setup des routes api
	router.GET("/api/albums", getAlbums)

	router.Run("localhost:8080")
}
