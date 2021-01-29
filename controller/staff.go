package controller

import (
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/staff/get?id=1
//后端通过c.Query("id")获取参数
func StaffGetOne(c *gin.Context) {
	var staff model.Staff
	id := c.Query("id")
	err := service.StaffService.GetOne(c, id, &staff)
	if err != nil {

	} else {
		Success(c, "GetOne", gin.H{
			"staff": staff,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/staff/get?id=1
//后端通过c.Query("id")获取参数
func StaffGetAll(c *gin.Context) {
	var staffes []model.Staff
	err := service.StaffService.GetAll(c, &staffes)
	if err != nil {

	} else {
		Success(c, "GetAll", gin.H{
			"staff": staffes,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/staff/first
func StaffGetFirst(c *gin.Context) {
	var staff model.Staff
	err := service.StaffService.First(c, &staff)
	if err != nil {

	} else {
		Success(c, "First", gin.H{
			"staff": staff,
		})
	}
}

//添加操作
//通过POST请求，前端将staff对象以post请求方式发送过来
//前端请求：/v2/staff/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到staff结构体中
func StaffNew(c *gin.Context) {
	var staff model.Staff
	c.ShouldBind(&staff)
	err := service.StaffService.New(c, &staff)
	if err != nil {

	} else {
		Success(c, "New", gin.H{
			"staff": staff,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/staff/delete?id=4&id=5
//后端获取id数组至ids
func StaffDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.StaffService.Delete(c, ids)
	if err != nil {

	} else {
		Success(c, "New", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/staff/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func StaffUpdate(c *gin.Context) {
	var staff model.Staff
	c.ShouldBind(&staff)
	err := service.StaffService.Update(c, &staff)
	if err != nil {

	} else {
		Success(c, "Update", gin.H{
			"staff": staff,
		})
	}

}
