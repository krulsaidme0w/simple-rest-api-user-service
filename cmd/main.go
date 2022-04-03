package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"golang_pet_project_1/internal/core/services/usersrv"
	"golang_pet_project_1/internal/handlers/userhandler"
	"golang_pet_project_1/internal/repositories/userrepo"
	"log"
	"sync"
)

const host = "0.0.0.0"
const port = "8080"

func main() {

	m := &sync.RWMutex{}

	userRepository := userrepo.Storage{
		Filename: "db.txt",
		Mutex:    m,
	}

	userService := usersrv.Service{
		Storage: userRepository,
	}

	userHandler := userhandler.UserHandler{
		UserService: userService,
	}

	r := router.New()
	r.POST("/user/", userHandler.CreateUser)
	r.GET("/user/get/", userHandler.GetUser)
	r.POST("/user/{user_id}", userHandler.UpdateUser)
	r.DELETE("/user/{user_id}", userHandler.DeleteUser)

	if err := fasthttp.ListenAndServe(host+":"+port, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
