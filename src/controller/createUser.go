package controller

import (
	"fmt"
	"net/http"

	"github.com/DevzIcaro/golang_crud/src/configuration/logger"
	"github.com/DevzIcaro/golang_crud/src/configuration/validation"
	"github.com/DevzIcaro/golang_crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Iniciação createuser controller",
		zap.String("Journey", "CreateString"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		logger.Error("Erro ao validar UserRequest", err,
			zap.String("Journey", "CreateString"),
		)

		restERR := validation.ValidateUserError(err)
		c.JSON(restERR.Code, restERR)
		return
	}

	logger.Info("UserRequest é valido",
		zap.String("Journey", "CreateString"),
	)

	userResponsePostman := request.UserRequest{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Name:     userRequest.Name,
		Age:      userRequest.Age,
	}

	c.JSON(http.StatusCreated, userResponsePostman)

	fmt.Println(userRequest)
}
