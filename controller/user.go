package controller

import (
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/user/get?id=1
//后端通过c.Query("id")获取参数
func UserGetOne(c *gin.Context) {
	var user model.User
	username := c.Query("username")
	err := service.UserService.GetOne(c, username, &user)
	if err != nil {

	} else {
		Success(c, "GetOne", gin.H{
			"user": user,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/user/get?id=1
//后端通过c.Query("id")获取参数
func UserGetAll(c *gin.Context) {
	var useres []model.User
	err := service.UserService.GetAll(c, &useres)
	if err != nil {

	} else {
		Success(c, "GetAll", gin.H{
			"user": useres,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/user/first
func UserGetFirst(c *gin.Context) {
	var user model.User
	err := service.UserService.First(c, &user)
	if err != nil {

	} else {
		Success(c, "First", gin.H{
			"user": user,
		})
	}
}

//添加操作
//通过POST请求，前端将user对象以post请求方式发送过来
//前端请求：/v2/user/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到user结构体中
func UserNew(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	err := service.UserService.New(c, &user)
	if err != nil {

	} else {
		Success(c, "New", gin.H{
			"user": user,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/user/delete?id=4&id=5
//后端获取id数组至ids
func UserDelete(c *gin.Context) {
	var ids []string
	usernames := c.QueryArray("username")
	err := service.UserService.Delete(c, usernames)
	if err != nil {

	} else {
		Success(c, "Delete", gin.H{
			"usernames": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/user/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func UserUpdate(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	err := service.UserService.Update(c, &user)
	if err != nil {

	} else {
		Success(c, "Update", gin.H{
			"user": user,
		})
	}

}
