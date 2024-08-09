package main

import (
	"fmt"
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

func getBalance(c *gin.Context) {
	address, balance, err := utils.GetBalance("0x2b754dEF498d4B6ADada538F01727Ddf67D91A7D")
	fmt.Println(err)
	c.JSON(200, gin.H{
		"address": address,
		"balance": balance,
	})

}

func main() {
	r := gin.Default()
	r.GET("/hello", sayHello)
	r.GET("/verifyMerkle", verifyMerkle)
	r.GET("/balance", getBalance)

	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
