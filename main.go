package main

import (
	"log"

	"github.com/DevzIcaro/golang_crud/src/configuration/logger"
	"github.com/DevzIcaro/golang_crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Iniciando a aplicação")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default() // rota com logger e middleware de recovery

	routes.InitRoutes(&router.RouterGroup) //Aqui pegamos o objeto original passando o ponteiro (&) e referenciando o RouterGroup que criei

	if err := router.Run(":8080"); err != nil { //se ja existir uma aplicação rodando nessa porta é retornado um erro
		log.Fatal(err)
	}

}
