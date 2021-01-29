package router

import (
	"CarNetBack/controller"
	"CarNetBack/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {
	router.Use(middleware.Cors())
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

		log := v1.Group("/log")
		{
			log.GET("/get", controller.LogGetOne)
			log.GET("/first", controller.LogGetFirst)
			log.GET("/all", controller.LogGetAll)
			log.POST("/new", controller.LogNew)
			log.DELETE("/delete", controller.LogDelete)
			log.PUT("/update", controller.LogUpdate)
		}

		permission := v1.Group("/permission")
		{
			permission.GET("/get", controller.PermissionGetOne)
			permission.GET("/first", controller.PermissionGetFirst)
			permission.GET("/all", controller.PermissionGetAll)
			permission.POST("/new", controller.PermissionNew)
			permission.DELETE("/delete", controller.PermissionDelete)
			permission.PUT("/update", controller.PermissionUpdate)
		}

		setting := v1.Group("/setting")
		{
			setting.GET("/get", controller.SettingGetOne)
			setting.GET("/first", controller.SettingGetFirst)
			setting.GET("/all", controller.SettingGetAll)
			setting.POST("/new", controller.SettingNew)
			setting.DELETE("/delete", controller.SettingDelete)
			setting.PUT("/update", controller.SettingUpdate)
		}

		staff := v1.Group("/staff")
		{
			staff.GET("/get", controller.StaffGetOne)
			staff.GET("/first", controller.StaffGetFirst)
			staff.GET("/all", controller.StaffGetAll)
			staff.POST("/new", controller.StaffNew)
			staff.DELETE("/delete", controller.StaffDelete)
			staff.PUT("/update", controller.StaffUpdate)
		}

		task := v1.Group("/task")
		{
			task.GET("/get", controller.TaskGetOne)
			task.GET("/first", controller.TaskGetFirst)
			task.GET("/all", controller.TaskGetAll)
			task.POST("/new", controller.TaskNew)
			task.DELETE("/delete", controller.TaskDelete)
			task.PUT("/update", controller.TaskUpdate)
		}

		user := v1.Group("/user")
		{
			user.GET("/get", controller.UserGetOne)
			user.GET("/first", controller.UserGetFirst)
			user.GET("/all", controller.UserGetAll)
			user.POST("/new", controller.UserNew)
			user.DELETE("/delete", controller.UserDelete)
			user.PUT("/update", controller.UserUpdate)
		}

		vehicle := v1.Group("/vehicle")
		{
			vehicle.GET("/get", controller.VehicleGetOne)
			vehicle.GET("/first", controller.VehicleGetFirst)
			vehicle.GET("/all", controller.VehicleGetAll)
			vehicle.POST("/new", controller.VehicleNew)
			vehicle.DELETE("/delete", controller.VehicleDelete)
			vehicle.PUT("/update", controller.VehicleUpdate)
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
