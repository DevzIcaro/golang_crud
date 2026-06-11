package controller

import (
	"fmt"

	rest_err "github.com/DevzIcaro/golang_crud/src/configuration/rest_error"
	"github.com/DevzIcaro/golang_crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restERR := rest_err.NewBadRequestError(
			fmt.Sprintf("Algum campo está incorreto, error=%s", err))

		c.JSON(restERR.Code, restERR)
		return
	}

	fmt.Println(userRequest)
}
