package routes

import (
	controller "github.com/DevzIcaro/golang_crud/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET(":/getUserById/:userId", controller.FindUser)
	r.GET(":/getUserByEmail/:userEmail", controller.FindEmail)
	r.POST(":/createUser", controller.CreateUser)
	r.PUT(":/updateUser/:userId", controller.UpdateUser)
	r.DELETE(":/deleteUser/:userId", controller.DeleteUser)

}
