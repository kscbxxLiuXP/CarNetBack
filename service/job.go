package service

import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type jobService struct {
}

// JobService create JobService instance
var JobService = jobService{}

func (us jobService) First(ctx *gin.Context, job *model.Job) error {
	return db.First(&job).Error
}

func (us jobService) GetOne(ctx *gin.Context, id string, job *model.Job) error {
	return db.Where("id = ?", id).Find(&job).Error
}

func (us jobService) GetAll(ctx *gin.Context, job *[]model.Job) error {
	return db.Find(&job).Error
}

func (us jobService) New(ctx *gin.Context, job *model.Job) error {
	return db.Create(&job).Error

}

func (us jobService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Job{}).Error
}

func (us jobService) Update(ctx *gin.Context, job *model.Job) error {
	return db.Save(&job).Error

}
