package user

import (
	_entities "Tugas/Project-Test-2/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetUser(id int) (_entities.User, int, error) {
	var users _entities.User
	tx := ur.database.Find(&users, id)
	if tx.Error != nil {
		return users, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, 0, tx.Error
	}
	return users, int(tx.RowsAffected), nil
}

func (ur *UserRepository) DeleteUser(id int) (_entities.User, error) {
	var users _entities.User
	tx := ur.database.Delete(&users, id)
	if tx.Error != nil {
		return users, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, tx.Error

	}
	return users, nil
}
func (ur *UserRepository) CreateUser(user _entities.User) (_entities.User, error) {
	tx := ur.database.Save(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, tx.Error

	}
	return user, nil
}
func (ur *UserRepository) UpdatedUser(users _entities.User) (_entities.User, error) {
	tx := ur.database.Save(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, tx.Error
	}
	return users, nil
}
