// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/golang-friends/members/internal/application"
	"github.com/golang-friends/members/internal/config"
	"github.com/golang-friends/members/internal/githubservice"
	"net/http"
)

// Injectors from wire.go:

func ProvideApplication() *application.Application {
	configConfig := config.FromViper()
	string2 := configConfig.Orgname
	client := _wireClientValue
	gitHubService := githubservice.New(string2, client)
	applicationApplication := application.NewApplication(configConfig, gitHubService)
	return applicationApplication
}

var (
	_wireClientValue = http.DefaultClient
)