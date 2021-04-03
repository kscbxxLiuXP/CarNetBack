package basic

import (
	"CarNetBack/controller"
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/job/get?id=1
//后端通过c.Query("id")获取参数
func JobGetOne(c *gin.Context) {
	var job model.Job
	id := c.Query("id")
	err := service.JobService.GetOne(c, id, &job)
	if err != nil {

	} else {
		controller.Success(c, "GetOne", gin.H{
			"job": job,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/job/get?id=1
//后端通过c.Query("id")获取参数
func JobGetAll(c *gin.Context) {
	var jobes []model.Job
	err := service.JobService.GetAll(c, &jobes)
	if err != nil {

	} else {
		controller.Success(c, "GetAll", gin.H{
			"job": jobes,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/job/first
func JobGetFirst(c *gin.Context) {
	var job model.Job
	err := service.JobService.First(c, &job)
	if err != nil {

	} else {
		controller.Success(c, "First", gin.H{
			"job": job,
		})
	}
}

//添加操作
//通过POST请求，前端将job对象以post请求方式发送过来
//前端请求：/v2/job/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到job结构体中
func JobNew(c *gin.Context) {
	var job model.Job
	c.ShouldBind(&job)
	err := service.JobService.New(c, &job)
	if err != nil {

	} else {
		controller.Success(c, "New", gin.H{
			"job": job,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/job/delete?id=4&id=5
//后端获取id数组至ids
func JobDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.JobService.Delete(c, ids)
	if err != nil {

	} else {
		controller.Success(c, "Delete", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/job/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func JobUpdate(c *gin.Context) {
	var job model.Job
	c.ShouldBind(&job)
	err := service.JobService.Update(c, &job)
	if err != nil {

	} else {
		controller.Success(c, "Update", gin.H{
			"job": job,
		})
	}

}

func JobCurrentJobByVehicleID(c *gin.Context) {
	var job model.Job
	vehicleID := c.Query("vehicleID")
	err := service.JobService.CurrentJobByVehicleID(c, vehicleID, &job)
	if err != nil {
		controller.Error(c, "error", gin.H{}, 1)
	} else {
		controller.Success(c, "GetOne", gin.H{
			"job": job,
		})
	}
}
func JobLastJobByVehicleID(c *gin.Context) {
	var job model.Job
	vehicleID := c.Query("vehicleID")
	err := service.JobService.LastJobByVehicleID(c, vehicleID, &job)
	if err != nil {
		controller.Error(c, "error", gin.H{}, 1)
	} else {
		controller.Success(c, "GetOne", gin.H{
			"job": job,
		})
	}
}
