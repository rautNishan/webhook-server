package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rautNishan/webhook-server.git/db"
)

type RegisterWebHookBody struct {
	URL string `json:"url"`
}

// This function returns a secret key
func RegisterWebHook(c *gin.Context) {
	var req RegisterWebHookBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("Received URL:", req.URL)

	c.JSON(200, gin.H{
		"secret_id": "time_random_string",
	})
}

func main() {
	router := gin.Default()
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	router.POST("/registerwebhook", RegisterWebHook)
	router.Run()
}
