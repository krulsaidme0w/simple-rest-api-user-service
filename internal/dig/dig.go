package dig

import (
	"log"

	"github.com/fasthttp/router"
	"go.uber.org/dig"

	"golang_pet_project_1/internal/config"
	"golang_pet_project_1/internal/core/ports"
	"golang_pet_project_1/internal/core/services/userservice"
	"golang_pet_project_1/internal/handlers/userhandler"
	userrepository "golang_pet_project_1/internal/repositories"
	"golang_pet_project_1/internal/repositories/cache"
	"golang_pet_project_1/internal/repositories/file"
	"golang_pet_project_1/pkg/errors/di_errors"
)

func GetContainer(storageType string) (*dig.Container, error) {
	if !(storageType == "fs" || storageType == "s3") {
		return nil, di_errors.BadStorageType
	}

	container := dig.New()
	if err := container.Provide(config.SetUp); err != nil {
		return nil, err
	}

	if storageType == "fs" {
		if err := container.Provide(file.NewStorage, dig.As(new(ports.UserStorage))); err != nil {
			return nil, err
		}
	}
	if storageType == "s3" {
		if err := container.Provide(file.NewMinioStorage, dig.As(new(ports.UserStorage))); err != nil {
			return nil, err
		}
	}

	if err := container.Provide(cache.NewCache); err != nil {
		return nil, err
	}

	if err := container.Invoke(fillCache); err != nil {
		return nil, err
	}

	if err := container.Provide(userrepository.NewStorage, dig.As(new(ports.UserRepository))); err != nil {
		return nil, err
	}

	if err := container.Provide(userservice.NewService, dig.As(new(ports.UserService))); err != nil {
		return nil, err
	}

	if err := container.Provide(userhandler.NewHandler); err != nil {
		return nil, err
	}

	if err := container.Provide(router.New); err != nil {
		return nil, err
	}

	if err := container.Invoke(initRouter); err != nil {
		return nil, err
	}

	return container, nil
}

func fillCache(storage ports.UserStorage, cache *cache.Cache) {
	users, err := storage.GetAllUsers()
	if err != nil {
		log.Fatal(err.Error())
	}
	cache.Fill(users)
}

func initRouter(r *router.Router, userHandler *userhandler.UserHandler) {
	r.POST("/user/", userHandler.CreateUser)
	r.GET("/user/get/", userHandler.GetUser)
	r.POST("/user/{user_id}/", userHandler.UpdateUser)
	r.DELETE("/user/{user_id}/", userHandler.DeleteUser)
}
