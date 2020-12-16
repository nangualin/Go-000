package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/go-kit/kit/log"
	"service"
)

func httpHandler(ctx context.Context,userService service.userService) Http.handler {
    r : mux.NewRouter()
    r.Methods("GET").Path('/info').Handler(userService.getUserInfo)
    return r
}
