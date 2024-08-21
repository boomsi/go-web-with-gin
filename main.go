package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/users", listUser)

	r.Run(":3000")
}

func listUser(c *gin.Context) {
	u := []user{
		{name: "abc", age: 18},
	}
	c.JSON(200, u)
}

type user struct {
	name string
	age  int
}
