package repositories

import (
	"github.com/ecommerce/internal/app/user/interfaces"
	"github.com/ecommerce/internal/model"
	"github.com/jinzhu/gorm"
)
type UserRepository interface{
	Create(user *model.User) error
	EmailExists(Email string)(bool,error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// func (r *UserRepositoryImpl) EmailExists(Email string) (bool, error) {
// 	var count int
// 	if err := r.db.Model(&model.User{}).Where("email = ?", Email).Count(&count).Error; err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }
// func (r *UserRepositoryImpl) Create(user *model.User) error {
// 	if err := r.db.Create(user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{DB}
}

