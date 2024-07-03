package repository

import "go-pzn-restful-api/model/domain"

type UserRepository interface {
	Save(user domain.User) domain.User
	Update(user domain.User) domain.User
	FindById(userId int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindByAddress(address string) (domain.User, error)
	Delete(userId int)
	DescBalanceUser(limit int) ([]domain.User, error)
	GetLastUserRank() (int, error)
	GetUserRank(userID int) (int, error)
}
