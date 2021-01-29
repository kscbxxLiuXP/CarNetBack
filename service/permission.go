package service
import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type permissionService struct {
}

// PermissionService create PermissionService instance
var PermissionService = permissionService{}

func (us permissionService) First(ctx *gin.Context, permission *model.Permission) error {
	return db.First(&permission).Error
}

func (us permissionService) GetOne(ctx *gin.Context, id string, permission *model.Permission) error {
	return db.Where("id = ?", id).Find(&permission).Error
}

func (us permissionService) GetAll(ctx *gin.Context, permission *[]model.Permission) error {
	return db.Find(&permission).Error
}

func (us permissionService) New(ctx *gin.Context, permission *model.Permission) error {
	return db.Create(&permission).Error

}

func (us permissionService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Permission{}).Error
}

func (us permissionService) Update(ctx *gin.Context, permission *model.Permission) error {
	return db.Save(&permission).Error

}
