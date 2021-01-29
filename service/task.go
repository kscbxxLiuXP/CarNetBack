package service
import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type taskService struct {
}

// TaskService create TaskService instance
var TaskService = taskService{}

func (us taskService) First(ctx *gin.Context, task *model.Task) error {
	return db.First(&task).Error
}

func (us taskService) GetOne(ctx *gin.Context, id string, task *model.Task) error {
	return db.Where("id = ?", id).Find(&task).Error
}

func (us taskService) GetAll(ctx *gin.Context, task *[]model.Task) error {
	return db.Find(&task).Error
}

func (us taskService) New(ctx *gin.Context, task *model.Task) error {
	return db.Create(&task).Error

}

func (us taskService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Task{}).Error
}

func (us taskService) Update(ctx *gin.Context, task *model.Task) error {
	return db.Save(&task).Error

}
