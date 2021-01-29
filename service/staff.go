package service
import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type staffService struct {
}

// StaffService create StaffService instance
var StaffService = staffService{}

func (us staffService) First(ctx *gin.Context, staff *model.Staff) error {
	return db.First(&staff).Error
}

func (us staffService) GetOne(ctx *gin.Context, id string, staff *model.Staff) error {
	return db.Where("id = ?", id).Find(&staff).Error
}

func (us staffService) GetAll(ctx *gin.Context, staff *[]model.Staff) error {
	return db.Find(&staff).Error
}

func (us staffService) New(ctx *gin.Context, staff *model.Staff) error {
	return db.Create(&staff).Error

}

func (us staffService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Staff{}).Error
}

func (us staffService) Update(ctx *gin.Context, staff *model.Staff) error {
	return db.Save(&staff).Error

}
