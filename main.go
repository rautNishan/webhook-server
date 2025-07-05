package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterWebHookBody struct {
	URL string `json:"url"`
}

// This function returns id that is needed to be passed while creating a task
func RegisterWebHook(c *gin.Context) {
	var req RegisterWebHookBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("Received URL:", req.URL)

	c.JSON(200, gin.H{
		"id": 1,
	})
}

func main() {
	router := gin.Default()

	router.POST("/registerwebhook", RegisterWebHook)

	router.Run()
}
