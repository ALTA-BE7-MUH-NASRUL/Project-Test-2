package user

import _entities "Tugas/Project-Test-2/entities"

type UserUseCaseInterface interface {
	GetUser(id int) (_entities.User, int, error)
	DeleteUser(id int) (_entities.User, error)
	CreateUser(user _entities.User) (_entities.User, error)
	UpdatedUser(users _entities.User, id int) (_entities.User, error)
}
