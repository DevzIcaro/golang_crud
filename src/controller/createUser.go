package controller

import (
	"fmt"
	"net/http"

	rest_err "github.com/DevzIcaro/golang_crud/src/configuration/rest_error"
	"github.com/DevzIcaro/golang_crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restERR := rest_err.NewBadRequestError(
			fmt.Sprintf("Algum campo está incorreto error=%s", err.Error))

		c.JSON(restERR.Code, restERR)
		return
	}

	userResponsePostman := request.UserRequest{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Name:     userRequest.Name,
		Age:      userRequest.Age,
	}

	c.JSON(http.StatusCreated, userResponsePostman)

	fmt.Println(userRequest)
}
