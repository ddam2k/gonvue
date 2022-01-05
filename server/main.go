package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"mgkim.com/gonvue/services/user"
)

func main() {

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	apiSvc := r.Group("api")
	{
		userSvc := apiSvc.Group("user")
		{
			userSvc.GET("list", user.List)
		}
	}

	r.Run("0.0.0.0:8082")
}
