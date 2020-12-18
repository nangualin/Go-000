package service

import (
    "errors"
    "fmt"
    "log"
    "net/http"
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
  GetUserInfo(w http.ResponseWriter, _ *http.Request)
}
type UserServiceImpl struct {
   message string
}

func UserServiceFun(msg string) UserService {
  return &UserServiceImpl{
    message:msg,
  }
}

func (userService *UserServiceImpl) GetUserInfo(w http.ResponseWriter, r *http.Request) {
  log.Printf("msg: %s", userService.message)
  fmt.Println("yes,you come in")
}
