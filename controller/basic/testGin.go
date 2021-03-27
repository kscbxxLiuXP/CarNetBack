//本代码作为CRUD操作基本实例模板
package basic

import (
	"CarNetBack/controller"
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/address/get?id=1
//后端通过c.Query("id")获取参数
func TestAddressGetOne(c *gin.Context) {
	var address model.Address
	id := c.Query("id")
	err := service.AddressService.GetOne(c, id, &address)
	if err != nil {

	} else {
		controller.Success(c, "GetOne", gin.H{
			"address": address,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/address/get?id=1
//后端通过c.Query("id")获取参数
func TestAddressGetAll(c *gin.Context) {
	var addresses []model.Address
	err := service.AddressService.GetAll(c, &addresses)
	if err != nil {

	} else {
		controller.Success(c, "GetAll", gin.H{
			"address": addresses,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/address/first
func TestAddressGetFirst(c *gin.Context) {
	var address model.Address
	err := service.AddressService.First(c, &address)
	if err != nil {

	} else {
		controller.Success(c, "First", gin.H{
			"address": address,
		})
	}
}

//添加操作
//通过POST请求，前端将address对象以post请求方式发送过来
//前端请求：/v2/address/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到Address结构体中
func TestAddressNew(c *gin.Context) {
	var address model.Address
	c.ShouldBind(&address)
	err := service.AddressService.New(c, &address)
	if err != nil {

	} else {
		controller.Success(c, "New", gin.H{
			"address": address,
		})
	}
}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/address/delete?id=4&id=5
//后端获取id数组至ids
func TestAddressDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.AddressService.Delete(c, ids)
	if err != nil {

	} else {
		controller.Success(c, "Delete", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/address/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func TestAddressUpdate(c *gin.Context) {
	var address model.Address
	c.ShouldBind(&address)
	err := service.AddressService.Update(c, &address)
	if err != nil {

	} else {
		controller.Success(c, "Update", gin.H{
			"address": address,
		})
	}

}
