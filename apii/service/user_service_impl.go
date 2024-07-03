package service

import (
	"errors"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository.UserRepository
	auth.JwtAuth
}

func (s *UserServiceImpl) Logout(userId int) web.UserResponse {
	findById, err := s.UserRepository.FindById(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	findById.Token = ""
	update := s.UserRepository.Update(findById)

	return helper.ToUserResponse(update)
}

func (s *UserServiceImpl) UploadAvatar(userId int, filePath string) web.UserResponse {
	findById, err := s.UserRepository.FindById(userId)
	oldAvatar := findById.Avatar
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if oldAvatar != filePath {
		if findById.Avatar == "" {
			return updateWhenUploadAvatar(findById, filePath, s.UserRepository)
		}
		os.Remove(oldAvatar)
		return updateWhenUploadAvatar(findById, filePath, s.UserRepository)
	}

	return updateWhenUploadAvatar(findById, filePath, s.UserRepository)
}

func (s *UserServiceImpl) FindById(userId int) web.UserResponse {
	findById, err := s.UserRepository.FindById(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	rankUser, err := s.UserRepository.GetUserRank(findById.Id)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	findById.Rank = rankUser
	update := s.UserRepository.Update(findById)

	return helper.ToUserResponse(update)
}

func (s *UserServiceImpl) FindDetailById(userId int) web.UserDetailResponse {
	findById, err := s.UserRepository.FindById(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	rankUser, err := s.UserRepository.GetUserRank(findById.Id)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	findById.Rank = rankUser
	update := s.UserRepository.Update(findById)

	lastRankUser, err := s.UserRepository.GetLastUserRank()
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToUserDetailResponse(update, lastRankUser)
}

func (s *UserServiceImpl) FindTop10User() []web.UserRankResponse {
	findTop10User, err := s.UserRepository.DescBalanceUser(10)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToUsersRankResponse(findTop10User)
}

func (s *UserServiceImpl) Login(input web.UserLoginInput) web.UserResponse {
	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	if err != nil {
		panic(helper.NewBadRequestError(errors.New("Email or password is wrong").Error()))
	}

	findByAddress, err := s.UserRepository.FindByAddress(input.Address)
	if findByAddress.Id == 0 || err != nil || findByAddress.Email != input.Email {
		panic(helper.NewBadRequestError(errors.New("Address is wrong").Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.Password), []byte(input.Password))
	if err != nil {
		panic(helper.NewBadRequestError(errors.New("Email or password is wrong").Error()))
	}

	token, _ := s.JwtAuth.GenerateJwtToken("user", findByEmail.Id)
	findByEmail.Token = token

	update := s.UserRepository.Update(findByEmail)

	return helper.ToUserResponse(update)
}

func (s *UserServiceImpl) Register(input web.UserRegisterInput) web.UserResponse {
	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	if findByEmail.Id != 0 || err != nil {
		panic(helper.NewNotFoundError(errors.New("email has been registered").Error()))
	}

	findByAddress, err := s.UserRepository.FindByAddress(input.Address)
	if findByAddress.Id != 0 || err != nil {
		panic(helper.NewNotFoundError(errors.New("Address has been registered").Error()))
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	domainUser := domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(password),
		Address:   input.Address,
	}
	save := s.UserRepository.Save(domainUser)
	helper.PanicIfError(err)

	auth := helper.AuthGenerator(helper.Client)
	add, err := helper.Manage.AddStudent(auth, common.HexToAddress(input.Address), big.NewInt(int64(save.Id)), input.LastName)
	if err != nil {
		helper.PanicIfError(err)
	}
	log.Printf("add successfull", add)

	return helper.ToUserResponse(save)
}

func NewUserService(userRepository repository.UserRepository, jwtAuth auth.JwtAuth) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		JwtAuth:        jwtAuth,
	}
}

func updateWhenUploadAvatar(user domain.User, filePath string, userRepository repository.UserRepository) web.UserResponse {
	user.Avatar = filePath
	update := userRepository.Update(user)
	response := helper.ToUserResponse(update)
	return response
}
