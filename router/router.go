package router

import (
	"CarNetBack/controller"
	"CarNetBack/controller/basic"
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
			address.GET("/get", basic.AddressGetOne)
			address.GET("/first", basic.AddressGetFirst)
			address.GET("/all", basic.AddressGetAll)

			address.POST("/new", basic.AddressNew)
			address.DELETE("/delete", basic.AddressDelete)
			address.PUT("/update", basic.AddressUpdate)
		}

		job := v1.Group("/job")
		{
			job.GET("/get", basic.JobGetOne)
			job.GET("/first", basic.JobGetFirst)
			job.GET("/all", basic.JobGetAll)

			job.POST("/new", basic.JobNew)
			job.DELETE("/delete", basic.JobDelete)
			job.PUT("/update", basic.JobUpdate)
		}

		log := v1.Group("/log")
		{
			log.GET("/get", basic.LogGetOne)
			log.GET("/first", basic.LogGetFirst)
			log.GET("/all", basic.LogGetAll)
			log.POST("/new", basic.LogNew)
			log.DELETE("/delete", basic.LogDelete)
			log.PUT("/update", basic.LogUpdate)
		}

		permission := v1.Group("/permission")
		{
			permission.GET("/get", basic.PermissionGetOne)
			permission.GET("/first", basic.PermissionGetFirst)
			permission.GET("/all", basic.PermissionGetAll)
			permission.POST("/new", basic.PermissionNew)
			permission.DELETE("/delete", basic.PermissionDelete)
			permission.PUT("/update", basic.PermissionUpdate)
		}

		setting := v1.Group("/setting")
		{
			setting.GET("/get", basic.SettingGetOne)
			setting.GET("/first", basic.SettingGetFirst)
			setting.GET("/all", basic.SettingGetAll)
			setting.POST("/new", basic.SettingNew)
			setting.DELETE("/delete", basic.SettingDelete)
			setting.PUT("/update", basic.SettingUpdate)
		}

		staff := v1.Group("/staff")
		{
			staff.GET("/get", basic.StaffGetOne)
			staff.GET("/first", basic.StaffGetFirst)
			staff.GET("/all", basic.StaffGetAll)
			staff.POST("/new", basic.StaffNew)
			staff.DELETE("/delete", basic.StaffDelete)
			staff.PUT("/update", basic.StaffUpdate)

			staff.GET("/checkForNew", basic.StaffCheckForNew)
			staff.GET("/getNextID", basic.StaffGetNextID)
			staff.POST("/uploadPhoto", basic.StaffUploadPhoto)
			staff.GET("/filterByCondition", basic.StaffFilterByCondition)
		}

		task := v1.Group("/task")
		{
			task.GET("/get", basic.TaskGetOne)
			task.GET("/first", basic.TaskGetFirst)
			task.GET("/all", basic.TaskGetAll)
			task.POST("/new", basic.TaskNew)
			task.DELETE("/delete", basic.TaskDelete)
			task.PUT("/update", basic.TaskUpdate)
		}

		user := v1.Group("/user")
		{
			user.GET("/get", basic.UserGetOne)
			user.GET("/first", basic.UserGetFirst)
			user.GET("/all", basic.UserGetAll)
			user.POST("/new", basic.UserNew)
			user.DELETE("/delete", basic.UserDelete)
			user.PUT("/update", basic.UserUpdate)
		}

		vehicle := v1.Group("/vehicle")
		{
			vehicle.GET("/get", basic.VehicleGetOne)
			vehicle.GET("/first", basic.VehicleGetFirst)
			vehicle.GET("/all", basic.VehicleGetAll)
			vehicle.POST("/new", basic.VehicleNew)
			vehicle.DELETE("/delete", basic.VehicleDelete)
			vehicle.PUT("/update", basic.VehicleUpdate)

			vehicle.POST("/registerInGroup", basic.VehicleRegisterInList)
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
			address.GET("/get", basic.TestAddressGetOne)
			address.GET("/first", basic.TestAddressGetFirst)
			address.GET("/all", basic.TestAddressGetAll)

			address.POST("/new", basic.TestAddressNew)
			address.DELETE("/delete", basic.TestAddressDelete)
			address.PUT("/update", basic.TestAddressUpdate)
		}

	}

	fmt.Println("[GIN-debug]", "Load routes successful.")
}
