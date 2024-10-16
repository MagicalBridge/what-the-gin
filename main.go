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

func generateMnemonic(c *gin.Context) {
	mnemonic, err := utils.GenerateMnemonic(256)
	fmt.Println(err)
	c.JSON(200, gin.H{
		"mnemonic": mnemonic,
	})
}

func generateP2PKHAddress(c *gin.Context) {
	p2pkh_Address := utils.GenerateP2PKHAddress()
	c.JSON(200, gin.H{
		"p2pkh_Address": p2pkh_Address,
	})
}

func generateP2SHAddress(c *gin.Context) {
	p2sh_Address, err := utils.GenerateP2SHAddress("02ef67f85c8376cf609a494af8c3a043df98211dec573cf1b0eb17304439cab90d")
	fmt.Println(err)

	c.JSON(200, gin.H{
		"p2sh_Address": p2sh_Address,
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", sayHello)
	r.GET("/verifyMerkle", verifyMerkle)
	r.GET("/balance", getBalance)
	r.GET("/generateMnemonic", generateMnemonic)
	r.GET("/generate_btc_p2pkh_address", generateP2PKHAddress)
	r.GET("/generate_btc_p2sh_address", generateP2SHAddress)

	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
