package main

import (
	"fmt"
	"log"
	"net/http"

	"gofinal/customers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("customer service")
	err := customers.CreateTable()
	if err != nil {
		log.Println(err)
		return
	}
	r := setRouter()
	r.Run(":2009")
	//run port ":2009"
}
func setRouter() *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth != "November 10, 2009" {
			c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			c.Abort()
			return
		}
		c.Next()
	})
	r.POST("/customers", customers.PostCreateCustomerHandler)
	r.GET("/customers/:id", customers.GetCustomerByIDHandler)
	r.GET("/customers", customers.GetAllCustomerHandler)
	r.PUT("/customers/:id", customers.PutUpdateCustomerHandler)
	r.DELETE("/customers/:id", customers.DeleteCustomerByIDHandler)
	return r
}
