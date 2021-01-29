package service
import (
	"CarNetBack/model"
	"github.com/gin-gonic/gin"
)

type userService struct {
}

// UserService create UserService instance
var UserService = userService{}

func (us userService) First(ctx *gin.Context, user *model.User) error {
	return db.First(&user).Error
}

func (us userService) GetOne(ctx *gin.Context, username string, user *model.User) error {
	return db.Where("username = ?", username).Find(&user).Error
}

func (us userService) GetAll(ctx *gin.Context, user *[]model.User) error {
	return db.Find(&user).Error
}

func (us userService) New(ctx *gin.Context, user *model.User) error {
	return db.Create(&user).Error

}

func (us userService) Delete(ctx *gin.Context, usernames []string) error {
	return db.Where("username in (?)", usernames).Delete(model.User{}).Error
}

func (us userService) Update(ctx *gin.Context, user *model.User) error {
	return db.Save(&user).Error

}
