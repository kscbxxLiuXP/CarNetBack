package basic

import (
	"CarNetBack/controller"
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/setting/get?id=1
//后端通过c.Query("id")获取参数
func SettingGetOne(c *gin.Context) {
	var setting model.Setting
	id := c.Query("id")
	err := service.SettingService.GetOne(c, id, &setting)
	if err != nil {

	} else {
		controller.Success(c, "GetOne", gin.H{
			"setting": setting,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/setting/get?id=1
//后端通过c.Query("id")获取参数
func SettingGetAll(c *gin.Context) {
	var settinges []model.Setting
	err := service.SettingService.GetAll(c, &settinges)
	if err != nil {

	} else {
		controller.Success(c, "GetAll", gin.H{
			"setting": settinges,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/setting/first
func SettingGetFirst(c *gin.Context) {
	var setting model.Setting
	err := service.SettingService.First(c, &setting)
	if err != nil {

	} else {
		controller.Success(c, "First", gin.H{
			"setting": setting,
		})
	}
}

//添加操作
//通过POST请求，前端将setting对象以post请求方式发送过来
//前端请求：/v2/setting/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到setting结构体中
func SettingNew(c *gin.Context) {
	var setting model.Setting
	c.ShouldBind(&setting)
	err := service.SettingService.New(c, &setting)
	if err != nil {

	} else {
		controller.Success(c, "New", gin.H{
			"setting": setting,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/setting/delete?id=4&id=5
//后端获取id数组至ids
func SettingDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.SettingService.Delete(c, ids)
	if err != nil {

	} else {
		controller.Success(c, "Delete", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/setting/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func SettingUpdate(c *gin.Context) {
	var setting model.Setting
	c.ShouldBind(&setting)
	err := service.SettingService.Update(c, &setting)
	if err != nil {

	} else {
		controller.Success(c, "Update", gin.H{
			"setting": setting,
		})
	}

}
