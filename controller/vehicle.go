package controller

import (
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/vehicle/get?id=1
//后端通过c.Query("id")获取参数
func VehicleGetOne(c *gin.Context) {
	var vehicle model.Vehicle
	id := c.Query("id")
	err := service.VehicleService.GetOne(c, id, &vehicle)
	if err != nil {

	} else {
		Success(c, "GetOne", gin.H{
			"vehicle": vehicle,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/vehicle/get?id=1
//后端通过c.Query("id")获取参数
func VehicleGetAll(c *gin.Context) {
	var vehiclees []model.Vehicle
	err := service.VehicleService.GetAll(c, &vehiclees)
	if err != nil {

	} else {
		Success(c, "GetAll", gin.H{
			"vehicle": vehiclees,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/vehicle/first
func VehicleGetFirst(c *gin.Context) {
	var vehicle model.Vehicle
	err := service.VehicleService.First(c, &vehicle)
	if err != nil {

	} else {
		Success(c, "First", gin.H{
			"vehicle": vehicle,
		})
	}
}

//添加操作
//通过POST请求，前端将vehicle对象以post请求方式发送过来
//前端请求：/v2/vehicle/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到vehicle结构体中
func VehicleNew(c *gin.Context) {
	var vehicle model.Vehicle
	c.ShouldBind(&vehicle)
	err := service.VehicleService.New(c, &vehicle)
	if err != nil {

	} else {
		Success(c, "New", gin.H{
			"vehicle": vehicle,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/vehicle/delete?id=4&id=5
//后端获取id数组至ids
func VehicleDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.VehicleService.Delete(c, ids)
	if err != nil {

	} else {
		Success(c, "New", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/vehicle/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func VehicleUpdate(c *gin.Context) {
	var vehicle model.Vehicle
	c.ShouldBind(&vehicle)
	err := service.VehicleService.Update(c, &vehicle)
	if err != nil {

	} else {
		Success(c, "Update", gin.H{
			"vehicle": vehicle,
		})
	}

}
