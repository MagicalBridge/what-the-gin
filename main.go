package main

import (
	"gin-web/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin",
	})
}

func verifyMerkle(c *gin.Context) {
	address := utils.VerifyMerkle()
	c.JSON(200, gin.H{
		"address": address,
	})
}

func main() {

	r := gin.Default()
	r.GET("/hello", sayHello)
	r.GET("/verifyMerkle", verifyMerkle)

	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
