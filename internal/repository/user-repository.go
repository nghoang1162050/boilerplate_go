package repository

import (
	"boilerplate_go/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User, roleId string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create implements UserRepository.
func (u *userRepository) Create(user *model.User, roleId string) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
            return err
        }

		userRole := model.UserRole{
            UserID: user.ID,
            RoleID:   roleId,
        }
        if err := tx.Create(&userRole).Error; err != nil {
            return err
        }

		return nil
	})
}
