package basic

import (
	"CarNetBack/controller"
	"CarNetBack/model"
	"CarNetBack/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

//查找操作：查找一个
//通过GET请求传递param
//前端请求：/v2/permission/get?id=1
//后端通过c.Query("id")获取参数
func PermissionGetOne(c *gin.Context) {
	var permission model.Permission
	id := c.Query("id")
	err := service.PermissionService.GetOne(c, id, &permission)
	if err != nil {

	} else {
		controller.Success(c, "GetOne", gin.H{
			"permission": permission,
		})
	}
}

//查找操作：查找全部
//通过GET请求传递param
//前端请求：/v2/permission/get?id=1
//后端通过c.Query("id")获取参数
func PermissionGetAll(c *gin.Context) {
	var permissiones []model.Permission
	err := service.PermissionService.GetAll(c, &permissiones)
	if err != nil {

	} else {
		controller.Success(c, "GetAll", gin.H{
			"permission": permissiones,
		})
	}
}

//查找操作：查找第一个
//不需要参数
//前端请求：/v2/permission/first
func PermissionGetFirst(c *gin.Context) {
	var permission model.Permission
	err := service.PermissionService.First(c, &permission)
	if err != nil {

	} else {
		controller.Success(c, "First", gin.H{
			"permission": permission,
		})
	}
}

//添加操作
//通过POST请求，前端将permission对象以post请求方式发送过来
//前端请求：/v2/permission/first
//后端：通过c.shouldBind，自动将前端表单数据绑定到permission结构体中
func PermissionNew(c *gin.Context) {
	var permission model.Permission
	c.ShouldBind(&permission)
	//检查权限是否已经存在
	var t model.CountStruct
	staffID := strconv.Itoa(permission.StaffID)
	vehicleID := strconv.Itoa(permission.VehicleID)
	service.PermissionService.HasPermission(c, staffID, vehicleID, &t)

	if t.Count <= 0 {
		service.PermissionService.New(c, &permission)
	}

	controller.Success(c, "New", gin.H{
		"permission": permission,
	})

}

//删除操作
//无论是删除单个还是删除多个都用一个函数解决
//前端请求：/v2/permission/delete?id=4&id=5
//后端获取id数组至ids
func PermissionDelete(c *gin.Context) {
	var ids []string
	ids = c.QueryArray("id")
	err := service.PermissionService.Delete(c, ids)
	if err != nil {

	} else {
		controller.Success(c, "Delete", gin.H{
			"ids": ids,
		})
	}

}

//更新操作
//使用UPDATE
//前端请求：/v2/permission/update?id=4&id=5
//后端将前端传递过来的对象进行Save操作，将新的值覆盖原始值
func PermissionUpdate(c *gin.Context) {
	var permission model.Permission
	c.ShouldBind(&permission)
	err := service.PermissionService.Update(c, &permission)
	if err != nil {

	} else {
		controller.Success(c, "Update", gin.H{
			"permission": permission,
		})
	}

}

func PermissionHasPermission(c *gin.Context) {
	var t model.CountStruct
	staffID := c.Query("staffID")
	vehicleID := c.Query("vehicleID")
	err := service.PermissionService.HasPermission(c, staffID, vehicleID, &t)
	var a int
	if t.Count > 0 {
		a = 1
	} else {
		a = 0
	}
	if err != nil {

	} else {
		controller.Success(c, "HasPermission", gin.H{
			"permission": a,
		})
	}
}

func PermissionGetByStaffID(c *gin.Context) {
	var permission []model.Permission
	staffID := c.Query("staffID")
	err := service.PermissionService.GetByStaffID(c, staffID, &permission)
	if err != nil {

	} else {
		controller.Success(c, "GetByStaffID", gin.H{
			"permission": permission,
		})
	}
}
func PermissionGetByVehicleID(c *gin.Context) {
	var permission []model.Permission
	vehicleID := c.Query("vehicleID")
	err := service.PermissionService.GetByVehicle(c, vehicleID, &permission)
	if err != nil {

	} else {
		controller.Success(c, "GetByVehicleID", gin.H{
			"permission": permission,
		})
	}
}
