package controller

import (
	"context"
	"net/http"
	"github.com/gorilla/mux"
	"Week04/service"
)

func HttpHandler(ctx context.Context,userService service.UserService) http.Handler {
    r := mux.NewRouter()
    r.HandleFunc("/info",userService.GetUserInfo)
    return r
}
