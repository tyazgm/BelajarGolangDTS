package service

import (
	"latihan/helper"
	"latihan/model"
	"latihan/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) Register(userRegisterRequest model.UserRegisterRequest) (string, error) {
	var id string
	id = helper.GenerateID()
	hashPassword, err := helper.Hash(userRegisterRequest.Password)
	if err != nil {
		return "", err
	}

	user := model.User{
		ID:       id,
		Email:    userRegisterRequest.Email,
		Password: hashPassword,
	}

	err = us.userRepository.Add(user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (us *UserService) Login(userLoginRequest model.UserLoginRequest) (string, error) {
	user, err := us.userRepository.GetByEmail(userLoginRequest.Email)
	if err != nil {
		return "", err
	}

	if !helper.IsHashValid(user.Password, userLoginRequest.Password) {
		return "", model.ErrorInvalidEmailOrPassword
	}

	token, err := helper.GenerateToken(user.ID)

	return token, err
}
