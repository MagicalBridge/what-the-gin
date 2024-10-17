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

func generate_btc_legacy_address(c *gin.Context) {
	publicKey := "02ef67f85c8376cf609a494af8c3a043df98211dec573cf1b0eb17304439cab90d"
	legacy_address, err := utils.Generate_btc_legacy_address(publicKey)
	fmt.Println(err)
	c.JSON(200, gin.H{
		"legacy_address": legacy_address, // 1CzkhKbqwmDL4o8StBdttLNesLpDZpddmA
	})
}

func generate_btc_nested_sigwit_address(c *gin.Context) {
	publicKey := "03156348ed9b36ea17115fa9eb05b58151847b8c96ce1ce78bd000cd620a0ca73c"
	nested_sigwit_address, err := utils.GenerateNestedSigwitddress(publicKey)
	fmt.Println(err)
	c.JSON(200, gin.H{
		"nested_sigwit_address": nested_sigwit_address, // 38s6gqg48tCkuGPSBzAcYphMkRofxf5M5K
	})
}

func generate_btc_native_sigwit_address(c *gin.Context) {
	publicKey := "02944695f65c4d602054f3260a0926a19b1f2941ffec043faa8144f60ccdef4646"
	nested_sigwit_address, err := utils.GenerateNativeSegWitAddress(publicKey)
	fmt.Println(err)
	c.JSON(200, gin.H{
		"nested_sigwit_address": nested_sigwit_address,
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", sayHello)
	r.GET("/verifyMerkle", verifyMerkle)
	r.GET("/balance", getBalance)
	r.GET("/generateMnemonic", generateMnemonic)
	r.GET("/generate_btc_legacy_address", generate_btc_legacy_address)
	r.GET("/generate_btc_nested_sigwit_address", generate_btc_nested_sigwit_address)
	r.GET("/generate_btc_native_sigwit_address", generate_btc_native_sigwit_address)

	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
