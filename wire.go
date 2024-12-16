//go:build wireinject
// +build wireinject

// wire.go
package main

import (
	"github.com/google/wire"
	"github.com/pealan/golang-rest-api/config"
	"github.com/pealan/golang-rest-api/handler"
	"github.com/pealan/golang-rest-api/repository"
	http "github.com/pealan/golang-rest-api/router"
)

func InitializeAPI() (*http.ServerHTTP, error) {
	wire.Build(config.ConnectToDB, repository.DeviceRepositoryInit, handler.DeviceHandlerInit, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
