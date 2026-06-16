package controller

import (
	"fmt"
	"net/http"

	"github.com/DevzIcaro/golang_crud/src/configuration/logger"
	"github.com/DevzIcaro/golang_crud/src/configuration/validation"
	"github.com/DevzIcaro/golang_crud/src/controller/model/request"
	"github.com/DevzIcaro/golang_crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.String(http.StatusCreated, "")

	fmt.Println(userRequest)
}
