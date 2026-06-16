package model

import (
	"github.com/DevzIcaro/golang_crud/src/configuration/logger"
	rest_err "github.com/DevzIcaro/golang_crud/src/configuration/rest_error"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}
