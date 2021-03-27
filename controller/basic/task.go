package basic

import (
	"CarNetBack/controller"
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/task/get?id=1
//后端通过c.Query("id")获取参数
func TaskGetOne(c *gin.Context) {
	var task model.Task
	id := c.Query("id")
	err := service.TaskService.GetOne(c, id, &task)
	if err != nil {

	} else {
		controller.Success(c, "GetOne", gin.H{
			"task": task,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/task/get?id=1
//后端通过c.Query("id")获取参数
func TaskGetAll(c *gin.Context) {
	var taskes []model.Task
	err := service.TaskService.GetAll(c, &taskes)
	if err != nil {

	} else {
		controller.Success(c, "GetAll", gin.H{
			"task": taskes,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/task/first
func TaskGetFirst(c *gin.Context) {
	var task model.Task
	err := service.TaskService.First(c, &task)
	if err != nil {

	} else {
		controller.Success(c, "First", gin.H{
			"task": task,
		})
	}
}

//添加操作
//通过POST请求，前端将task对象以post请求方式发送过来
//前端请求：/v2/task/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到task结构体中
func TaskNew(c *gin.Context) {
	var task model.Task
	c.ShouldBind(&task)
	err := service.TaskService.New(c, &task)
	if err != nil {

	} else {
		controller.Success(c, "New", gin.H{
			"task": task,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/task/delete?id=4&id=5
//后端获取id数组至ids
func TaskDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.TaskService.Delete(c, ids)
	if err != nil {

	} else {
		controller.Success(c, "Delete", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/task/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func TaskUpdate(c *gin.Context) {
	var task model.Task
	c.ShouldBind(&task)
	err := service.TaskService.Update(c, &task)
	if err != nil {

	} else {
		controller.Success(c, "Update", gin.H{
			"task": task,
		})
	}

}
