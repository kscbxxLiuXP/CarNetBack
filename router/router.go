package router

import (
	"CarNetBack/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {

	//404错误
	router.NoRoute(controller.NoRoute)

	v1 := router.Group("/v1")

	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v1 pong",
			})
		})
		address := v1.Group("/address")
		{
			address.GET("/get", controller.AddressGetOne)
			address.GET("/first", controller.AddressGetFirst)
			address.GET("/all", controller.AddressGetAll)

			address.POST("/new", controller.AddressNew)
			address.DELETE("/delete", controller.AddressDelete)
			address.PUT("/update", controller.AddressUpdate)
		}

		job := v1.Group("/job")
		{
			job.GET("/get", controller.JobGetOne)
			job.GET("/first", controller.JobGetFirst)
			job.GET("/all", controller.JobGetAll)

			job.POST("/new", controller.JobNew)
			job.DELETE("/delete", controller.JobDelete)
			job.PUT("/update", controller.JobUpdate)
		}
	}

	//测试golang数据模型与数据库属性绑定
	v2 := router.Group("/v2")
	{
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		address := v2.Group("/address")
		{
			address.GET("/get", controller.TestAddressGetOne)
			address.GET("/first", controller.TestAddressGetFirst)
			address.GET("/all", controller.TestAddressGetAll)

			address.POST("/new", controller.TestAddressNew)
			address.DELETE("/delete", controller.TestAddressDelete)
			address.PUT("/update", controller.TestAddressUpdate)
		}

	}

	fmt.Println("[GIN-debug]", "Load routes successful.")
}
