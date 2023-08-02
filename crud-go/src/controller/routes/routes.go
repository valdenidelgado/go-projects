package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId")
	r.GET("/getUserEmail/:userEmail")
	r.POST("/createUser")
	r.PUT("/updateUser/:userId")
	r.DELETE("/deleteUser/:userId")
}
