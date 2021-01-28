package main

import (
	"CarNetBack/config"
	"CarNetBack/model"
	"CarNetBack/router"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

func init() {

}

var address model.Address

func main() {
	r := gin.Default()

	config.LoadConfig()
	service.SetUp()
	router.LoadRoutes(r)

	r.Run(":8080")

}
