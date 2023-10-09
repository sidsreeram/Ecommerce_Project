package server

import (
	"net/http"

	"github.com/ecommerce/configure"
	"github.com/ecommerce/internal/app/admin/repositories"
	"github.com/ecommerce/internal/database"
)
func InitializeAPI(cfg configure.Config) (*http.ServerHTTP, error) {
	gormDB, err := database.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repositories.NewUserRepostiory
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}