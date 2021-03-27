package basic

import (
	"CarNetBack/controller"
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/log/get?id=1
//后端通过c.Query("id")获取参数
func LogGetOne(c *gin.Context) {
	var log model.Log
	id := c.Query("id")
	err := service.LogService.GetOne(c, id, &log)
	if err != nil {

	} else {
		controller.Success(c, "GetOne", gin.H{
			"log": log,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/log/get?id=1
//后端通过c.Query("id")获取参数
func LogGetAll(c *gin.Context) {
	var loges []model.Log
	err := service.LogService.GetAll(c, &loges)
	if err != nil {

	} else {
		controller.Success(c, "GetAll", gin.H{
			"log": loges,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/log/first
func LogGetFirst(c *gin.Context) {
	var log model.Log
	err := service.LogService.First(c, &log)
	if err != nil {

	} else {
		controller.Success(c, "First", gin.H{
			"log": log,
		})
	}
}

//添加操作
//通过POST请求，前端将log对象以post请求方式发送过来
//前端请求：/v2/log/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到log结构体中
func LogNew(c *gin.Context) {
	var log model.Log
	c.ShouldBind(&log)
	err := service.LogService.New(c, &log)
	if err != nil {

	} else {
		controller.Success(c, "New", gin.H{
			"log": log,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/log/delete?id=4&id=5
//后端获取id数组至ids
func LogDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.LogService.Delete(c, ids)
	if err != nil {

	} else {
		controller.Success(c, "Delete", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/log/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func LogUpdate(c *gin.Context) {
	var log model.Log
	c.ShouldBind(&log)
	err := service.LogService.Update(c, &log)
	if err != nil {

	} else {
		controller.Success(c, "Update", gin.H{
			"log": log,
		})
	}

}
