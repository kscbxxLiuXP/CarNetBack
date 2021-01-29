package service

import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type settingService struct {
}

// SettingService create SettingService instance
var SettingService = settingService{}

func (us settingService) First(ctx *gin.Context, setting *model.Setting) error {
	return db.First(&setting).Error
}

func (us settingService) GetOne(ctx *gin.Context, id string, setting *model.Setting) error {
	return db.Where("id = ?", id).Find(&setting).Error
}

func (us settingService) GetAll(ctx *gin.Context, setting *[]model.Setting) error {
	return db.Find(&setting).Error
}

func (us settingService) New(ctx *gin.Context, setting *model.Setting) error {
	return db.Create(&setting).Error

}

func (us settingService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Setting{}).Error
}

func (us settingService) Update(ctx *gin.Context, setting *model.Setting) error {
	return db.Save(&setting).Error

}
