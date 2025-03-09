package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var admin1 = User{ID: 1, Username: "BNK4970"}
var user1 = User{ID: 2, Username: "Lux"}

var rooms = []Room{
	{ID: 1, Label: "Chill", Description: "Un espace détente pour écouter de la musique, lire ou simplement se relaxer.", Users: []User{admin1}},
	{ID: 2, Label: "Gaming", Description: "Rejoins une partie endiablée avec tes amis et affrontez-vous dans vos jeux préférés.", Users: []User{user1}},
	{ID: 3, Label: "Devoir", Description: "Un endroit calme pour travailler efficacement et avancer sur tes projets scolaires.", Users: []User{}},
	{ID: 4, Label: "Débat", Description: "Exprime tes idées, écoute les arguments des autres et participe à des discussions passionnantes.", Users: []User{}},
}

func getRooms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rooms)
}

func getRoomByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	for _, room := range rooms {
		if room.ID == id {
			c.IndentedJSON(http.StatusOK, room)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
}

func main() {
	router := gin.Default()
	router.GET("/api/rooms", getRooms)
	router.GET("/api/rooms/:id", getRoomByID)
	router.Run("localhost:8080")
}
