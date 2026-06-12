package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevzIcaro/golang_crud/src/configuration/validation"
	"github.com/DevzIcaro/golang_crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Iniciação createuser controller")

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		log.Println("Iniciação createuser controller, error=", err.Error())

		restERR := validation.ValidateUserError(err)
		c.JSON(restERR.Code, restERR)
		return
	}

	log.Println("UserRequest é valido", userRequest)

	userResponsePostman := request.UserRequest{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Name:     userRequest.Name,
		Age:      userRequest.Age,
	}

	c.JSON(http.StatusCreated, userResponsePostman)

	fmt.Println(userRequest)
}
