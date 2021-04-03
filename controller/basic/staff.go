package basic

import (
	"CarNetBack/controller"
	"CarNetBack/model"
	"CarNetBack/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
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
		controller.Success(c, "GetOne", gin.H{
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
		controller.Success(c, "GetAll", gin.H{
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
		controller.Success(c, "First", gin.H{
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
		controller.Success(c, "New", gin.H{
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
		controller.Success(c, "Delete", gin.H{
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
		controller.Success(c, "Update", gin.H{
			"staff": staff,
		})
	}

}

//检查注册合法性
func StaffCheckForNew(c *gin.Context) {
	type T struct {
		ID       string `json:"id" form:"id"`
		IDNumber string `gorm:"column:idNumber" json:"idNumber" form:"idNumber"`
	}
	var t T
	var idEr, idNumEr = false, false
	c.ShouldBind(&t)
	//检查ID重复
	var staff model.Staff
	_ = service.StaffService.CheckForNewID(c, t.ID, &staff)
	if staff != (model.Staff{}) { //staff非空，已存在
		idEr = true
	}
	//清空staff，检查IDNumber重复
	staff = model.Staff{}
	_ = service.StaffService.CheckForNewIDNumber(c, t.IDNumber, &staff)
	if staff != (model.Staff{}) { //staff非空，已存在
		idNumEr = true
	}
	var err []string
	if idNumEr {
		err = append(err, "身份证号码已注册！")
	}
	if idEr {
		err = append(err, "ID已注册！")
	}

	if len(err) == 0 {
		//没有问题
		controller.Success(c, "StaffCheckForNew", gin.H{
			"staff": staff,
		})
	} else {
		controller.Error(c, "错误", gin.H{
			"error": err,
		}, 1)
	}
}

func StaffGetNextID(c *gin.Context) {

	var t model.MaxStruct
	err := service.StaffService.GetNextID(c, &t)
	if err != nil {

	} else {
		controller.Success(c, "StaffGetNextID", gin.H{
			"id": t.Max,
		})
	}

}

func StaffUploadPhoto(c *gin.Context) {

	//上传文件
	f, err := c.FormFile("file")
	id, _ := c.GetPostForm("id")
	//fmt.Println(id)
	if err != nil {
		controller.Success(c, "New", gin.H{
			"msg": "fail",
		})
	} else {
		//保存的文件
		//filePath := fmt.Sprintf("./%s", f.Filename)
		fileType := strings.Split(f.Filename, ".")[1]
		dst := path.Join("./data/avatar/", id+"."+fileType)
		fmt.Println(dst)

		_ = c.SaveUploadedFile(f, dst)

		var staff model.Staff
		_ = service.StaffService.GetOne(c, id, &staff)
		staff.Photoed = 1
		_ = service.StaffService.Update(c, &staff)

		controller.Success(c, "New", gin.H{
			"msg": "success",
		})
	}

}


func StaffFilterByCondition(c *gin.Context) {
	var staff []model.Staff
	id := c.DefaultQuery("id", "")
	name := c.DefaultQuery("name", "")
	idNumber := c.DefaultQuery("idNumber", "")
	err := service.StaffService.FilterByCondition(c, id, name, idNumber, &staff)
	if err != nil {

	} else {
		controller.Success(c, "GetOne", gin.H{
			"staff": staff,
		})
	}
}

func StaffAvatar(c *gin.Context) {
	id := c.Query("id")
	var p = ""
	p = path.Join("./data/avatar/", id+".png")
	c.File(p)
}
