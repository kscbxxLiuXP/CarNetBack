package service

import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type vehicleService struct {
}

// VehicleService create VehicleService instance
var VehicleService = vehicleService{}

func (us vehicleService) First(ctx *gin.Context, vehicle *model.Vehicle) error {
	return db.First(&vehicle).Error
}

func (us vehicleService) GetOne(ctx *gin.Context, id string, vehicle *model.Vehicle) error {
	return db.Where("id = ?", id).Find(&vehicle).Error
}

func (us vehicleService) GetAll(ctx *gin.Context, vehicle *[]model.Vehicle) error {
	return db.Find(&vehicle).Error
}

func (us vehicleService) New(ctx *gin.Context, vehicle *model.Vehicle) error {
	return db.Create(&vehicle).Error

}

func (us vehicleService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Vehicle{}).Error
}

func (us vehicleService) Update(ctx *gin.Context, vehicle *model.Vehicle) error {
	return db.Save(&vehicle).Error

}
func (us vehicleService) FilterByCondition(ctx *gin.Context, address string, name string, state string, vehicle *[]model.Vehicle) error {
	return db.Where("addressID like '%" + address + "%'").Where("state like '%" + state + "%'").Where("name like '%" + name + "%'").Find(&vehicle).Error
}
