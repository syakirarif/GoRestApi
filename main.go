package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi/auth"
)

func main() {
	route := gin.Default()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"username": "syakirarif",
		})
	})
	route.GET("/profile", auth.Profile)
	route.GET("/register", auth.Register)
	route.GET("/registerform", auth.RegisterForm)
	route.GET("profile/:username", auth.Category)
	route.GET("/showprofile", auth.ShowUser)
	route.GET("/showposting", auth.ShowPosting)
	route.GET("/detailprofile", auth.ShowDetailUser)

	// create user from raw
	route.POST("/createprofile", auth.CreateUser)

	//migration
	route.GET("migratedb", auth.MigrateTable)

	//http client
	route.GET("/grabuser", auth.GrabUser)

	route.Run(":8989")

	//dashboard := route.Group("/dashboard")
	//dashboard.Use()
}
