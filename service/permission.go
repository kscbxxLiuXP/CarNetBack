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

func (us permissionService) GetByStaffID(ctx *gin.Context, staffID string, permission *[]model.Permission) error {
	return db.Where("staffID = ?", staffID).Find(&permission).Error
}

func (us permissionService) GetByVehicle(ctx *gin.Context, vehicleID string, permission *[]model.Permission) error {
	return db.Where("vehicleID = ?", vehicleID).Find(&permission).Error
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

func (us permissionService) HasPermission(ctx *gin.Context, staffID string, vehicleID string, t *model.CountStruct) error {
	return db.Raw("SELECT COUNT(*) AS count FROM permission Where staffID=" + staffID + " and vehicleID=" + vehicleID).Scan(&t).Error

}
