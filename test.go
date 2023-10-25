package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Name    string 	`json:"name"`
	Age     int 	`json:"age"`
	Address string 	`json:"address"`
	Phone   string 	`json:"phone"`
}

type Transaction struct {
	Nama          string `json:"nama"`
	PaymentMethod string `json:"paymentMethod"`
	PaymentAmount string `json:"paymentAmount"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

func register(c *gin.Context) {
	var register User
	err := c.ShouldBindJSON(&register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"data": register,
		})
	}
}

func getAllUser(c *gin.Context) {
	var users []User
	c.JSON(http.StatusOK, users)
}

func login(c *gin.Context) {
	var userCredent UserCredential
	err := c.ShouldBindJSON(&userCredent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"username": userCredent.Username,
		})
	}
}

func createTransaction(c *gin.Context) {
	var transaction Transaction
	err := c.ShouldBindJSON(&userCredent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	c.JSON(http.StatusOK, transaction)
	}
}

func getAllTransaction(c *gin.Context) {
	var transactions []Transaction
	c.JSON(http.StatusOK, transactions)
}

func GetTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction Transaction
	c.JSON(http.StatusOK, transaction)
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction deleted successfully",
	})
}

func main() {
	fmt.Println("Hello, World!")
	r := gin.New()

	r.POST("/register", register)
	r.GET("/alluser", getAllUser)
	r.POST("/login", login)

	r.POST("/transaction", createTransaction)
	r.GET("/transaction", getAllTransaction)
	r.GET("/transaction/:id", GetTransaction)
	r.PUT("/transaction/:id", UpdateTransaction)
	r.DELETE("/transaction/:id", DeleteTransaction)

	r.Run()
}
