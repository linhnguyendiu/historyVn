package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := r.db.Find(&user, "email=?", email).Error
	if err != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByAddress(address string) (domain.User, error) {
	user := domain.User{}
	err := r.db.Find(&user, "address=?", address).Error
	if err != nil {
		return user, errors.New("Address is not belong to you")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Save(user domain.User) domain.User {
	err := r.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) Update(user domain.User) domain.User {
	err := r.db.Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) FindById(userId int) (domain.User, error) {
	user := domain.User{}
	//err := r.db.Preload("Courses").Where("id=?", userId).Find(&user).Error
	err := r.db.Where("id=?", userId).Find(&user).Error
	if err != nil || user.Id == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(userId int) {
	user := domain.User{}
	err := r.db.Where("id=?", userId).Delete(&user).Error

	helper.PanicIfError(err)
}

func (r *UserRepositoryImpl) DescBalanceUser(limit int) ([]domain.User, error) {
	users := []domain.User{}
	if err := r.db.Table("users").
		Order("balance DESC").
		Find(&users).Error; err != nil {
		return nil, err
	}

	// Cập nhật rank
	currentRank := 1
	sameRankCount := 0
	var currentBalance *int

	for i := range users {
		if currentBalance == nil || users[i].Balance != *currentBalance {
			currentRank += sameRankCount
			sameRankCount = 1
		} else {
			sameRankCount++
		}
		users[i].Rank = currentRank
		currentBalance = &users[i].Balance
	}

	// Lọc danh sách người dùng theo thứ hạng từ 1 đến limit
	topRankedUsers := []domain.User{}
	for _, user := range users {
		if user.Rank <= limit {
			topRankedUsers = append(topRankedUsers, user)
		}
	}

	// Cập nhật thứ hạng trong database
	for _, user := range topRankedUsers {
		if err := r.db.Table("users").Where("id = ?", user.Id).Update("rank", user.Rank).Error; err != nil {
			return nil, err
		}
	}

	return topRankedUsers, nil
}

func (r *UserRepositoryImpl) CalculateUserRanks() ([]domain.User, error) {
	users := []domain.User{}
	if err := r.db.Table("users").
		Order("balance DESC").
		Find(&users).Error; err != nil {
		return nil, err
	}

	// Cập nhật rank
	currentRank := 1
	sameRankCount := 0
	var currentBalance *int

	for i := range users {
		if currentBalance == nil || users[i].Balance != *currentBalance {
			currentRank += sameRankCount
			sameRankCount = 1
		} else {
			sameRankCount++
		}
		users[i].Rank = currentRank
		currentBalance = &users[i].Balance
	}

	// Lọc danh sách người dùng theo thứ hạng từ 1 đến limit
	topRankedUsers := []domain.User{}
	for _, user := range users {
		topRankedUsers = append(topRankedUsers, user)
	}

	// Cập nhật thứ hạng trong database
	for _, user := range topRankedUsers {
		if err := r.db.Table("users").Where("id = ?", user.Id).Update("rank", user.Rank).Error; err != nil {
			return nil, err
		}
	}

	return topRankedUsers, nil
}

// Hàm lấy thứ hạng của một người dùng cụ thể
func (r *UserRepositoryImpl) GetUserRank(userID int) (int, error) {
	users, err := r.CalculateUserRanks()
	if err != nil {
		return 0, err
	}

	for _, user := range users {
		if user.Id == userID {
			return user.Rank, nil
		}
	}

	return 0, errors.New("user with ID not found")
}

// Hàm lấy thứ hạng của người dùng đứng cuối cùng
func (r *UserRepositoryImpl) GetLastUserRank() (int, error) {
	users, err := r.CalculateUserRanks()
	if err != nil {
		return 0, err
	}

	if len(users) == 0 {
		return 0, errors.New("no users found")
	}

	return users[len(users)-1].Rank, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
