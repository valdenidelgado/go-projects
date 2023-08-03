package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
