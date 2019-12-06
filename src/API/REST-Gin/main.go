package main

import (
	"log"

	// "github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", Get)
	route.POST("/testing", Post)

	route.Run(":8080")
}

func Get(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

	c.String(200, "Success")
}

func Post(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

	c.String(200, "Success")
}
