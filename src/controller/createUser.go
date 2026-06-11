package controller

import (
	rest_err "github.com/DevzIcaro/golang_crud/src/configuration/rest_error"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("A criação de usuário é inválida") // mensagem exlcusiva da rota ao ser chamada, que puxa a formatação do NewBadRequestError
	c.JSON(err.Code, err)
}
