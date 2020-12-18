//+build wireinject

package main

import (
    "github.com/google/wire"
    "github.com/gorilla/mux"
    "context"
    "net/http"
    "Week04/service"
    "Week04/controller"
)

func InitHttpHandler(msg string ,ctx context.Context) http.Handler {
    wire.Build(service.UserServiceFun,controller.HttpHandler)
    return &mux.Router{}
}