package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"imooc-product/datamodels"
	"imooc-product/repositories"
)

type IUserService interface {
	IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOK bool)
	AddUser(user *datamodels.User) (userId int64, err error)
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository}
}

func (u *UserService) IsPwdSuccess(username string, pwd string) (user *datamodels.User, isOK bool) {
	var err error
	user, err = u.UserRepository.Select(username)
	if err != nil {
		return
	}
	isOK, _ = ValidatePassword(pwd, user.HashPassword)
	if !isOK {
		return &datamodels.User{}, false
	}
	return
}

//密码进行默认加密
func GeneratePassword(userPAssword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPAssword), bcrypt.DefaultCost)
}

//匹配密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码不匹配")
	}
	return true, nil
}

func (u *UserService) AddUser(user *datamodels.User) (userId int64, err error) {
	passWordHash, errPwd := GeneratePassword(user.HashPassword)
	if errPwd != nil {
		return userId, errPwd
	}
	user.HashPassword = string(passWordHash)
	return u.UserRepository.Insert(user)
}
