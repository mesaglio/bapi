package main

import (
	"golang-test/api/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ping := r.Group("/ping")
	{
		ping.GET("", controllers.Ping)
	}

	users := r.Group("/users")
	{
		users.GET("", controllers.GetUsers)
		users.POST("", controllers.AddUser)
		users.GET(":username", controllers.GetUserByUsername)
		users.DELETE(":username", controllers.DeleteUserByUsername)
		users.PATCH(":username", controllers.UpdateUserByUsername)
	}

	r.Run(":8080")
}
