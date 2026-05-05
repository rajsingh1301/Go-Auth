package main

import (
	"auth-go-test/controllers"
	"auth-go-test/initializers"
	"auth-go-test/middleware"

	"github.com/gin-gonic/gin"
)


func init(){
	initializers.LoadEnv() // Loading environment variables
	initializers.ConnectToDB() // Connecting to database
	initializers.SyncDatabase() // Syncing model schema to database(Auto Migrating)
	 
	
}
func main() {

	 router := gin.Default()
  router.POST("/signup", controllers.SignUp)
  router.POST("/login", controllers.Login)
  router.GET("/validate", middleware.RequireAuth, controllers.Validate)

  router.Run() // listens on 0.0.0.0:8080 by default

}