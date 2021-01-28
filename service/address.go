package service

import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type addressService struct {
}

// AddressService create addressService instance
var AddressService = addressService{}

func (us addressService) First(ctx *gin.Context, address *model.Address) error {
	return db.First(&address).Error
}

func (us addressService) GetOne(ctx *gin.Context, id string, address *model.Address) error {
	return db.Where("id = ?", id).Find(&address).Error
}

func (us addressService) GetAll(ctx *gin.Context, address *[]model.Address) error {
	return db.Find(&address).Error
}

func (us addressService) New(ctx *gin.Context, address *model.Address) error {
	return db.Create(&address).Error

}

func (us addressService) Delete(ctx *gin.Context, ids []string) error {
	return db.Where("id in (?)", ids).Delete(model.Address{}).Error
}

func (us addressService) Update(ctx *gin.Context, address *model.Address) error {
	return db.Save(&address).Error

}
