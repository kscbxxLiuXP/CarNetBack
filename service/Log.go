package service
import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type logService struct {
}

// LogService create LogService instance
var LogService = logService{}

func (us logService) First(ctx *gin.Context, log *model.Log) error {
	return db.First(&log).Error
}

func (us logService) GetOne(ctx *gin.Context, id string, log *model.Log) error {
	return db.Where("id = ?", id).Find(&log).Error
}

func (us logService) GetAll(ctx *gin.Context, log *[]model.Log) error {
	return db.Find(&log).Error
}

func (us logService) New(ctx *gin.Context, log *model.Log) error {
	return db.Create(&log).Error

}

func (us logService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Log{}).Error
}

func (us logService) Update(ctx *gin.Context, log *model.Log) error {
	return db.Save(&log).Error

}
