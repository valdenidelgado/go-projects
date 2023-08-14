package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", model.VerifyTokenMiddleware, userController.CreateUser)
	r.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
}
