// +build wireinject

package main

import (
    "github.com/google/wire"
    "github.com/gorilla/mux"
    "context"
    "net/http"
    "service"

)

func InitHttpHandler(msg String ,ctx context.Context) http.Handler {
    wire.Build(service.userService,controller.httpHandler)
    return &mux.Router{}
}