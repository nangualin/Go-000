package service

import (
  "context"
  "errors"
  "log"
)
type UserInfoDTO struct {
   ID int64
   Username string
   Age int
}
var (
  ErrUserNotExisted = errors.New("user is not existed")
  ErrPassword = errors.New("username or password are not match")
)

type UserService interface {
  getUserInfo(ctx context.Context,userName,password string) (*UserInfoDTO ,error)
}
type UserServiceImpl struct {
   message string
}

func UserServiceFun(msg string) UserService {
  return &UserServiceImpl{
    message:msg,
  }
}

func (userService *UserServiceImpl) getUserInfo(ctx context.Context,userName,password string) (*UserInfoDTO ,error) {
  err := ErrUserNotExisted
  log.Printf("err: %s", userService.message)
  return &UserInfoDTO{1,userName,12} ,err
}
