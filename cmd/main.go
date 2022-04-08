package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fasthttp/router"
	iradix "github.com/hashicorp/go-immutable-radix"
	"github.com/valyala/fasthttp"

	"golang_pet_project_1/internal/config"
	"golang_pet_project_1/internal/core/services/userservice"
	"golang_pet_project_1/internal/handlers/userhandler"
	userrepository "golang_pet_project_1/internal/repositories"
	usercache "golang_pet_project_1/internal/repositories/cache"
	"golang_pet_project_1/internal/repositories/file"
)

func main() {
	c, err := config.SetUp()
	c.Port = "8080"
	c.Host = "0.0.0.0"
	c.DB = "../db"

	//if err != nil {
	//	c.Port = "8080"
	//	c.Host = "0.0.0.0"
	//	c.DB = "db"
	//	//log.Fatal(err.Error())
	//}

	fileMutex := &sync.RWMutex{}
	prefixMutex := &sync.RWMutex{}

	usernamePrefix := iradix.New()
	namePrefix := iradix.New()

	cache := usercache.NewCache(usernamePrefix, namePrefix, prefixMutex)

	start := time.Now()
	err = cache.FillFromDB(c.DB)
	if err != nil {
		log.Fatalf(err.Error())
	}
	end := time.Now()

	fmt.Println("FILLED FROM DB: ", end.Sub(start).Seconds(), "s")

	fsStorage := file.NewStorage(c.DB, fileMutex)

	userRepository := userrepository.NewStorage(fsStorage, cache)
	userService := userservice.NewService(userRepository)
	userHandler := userhandler.NewHandler(userService)

	r := router.New()
	r.POST("/user/", userHandler.CreateUser)
	r.GET("/user/get/", userHandler.GetUser)
	r.POST("/user/{user_id}", userHandler.UpdateUser)
	r.DELETE("/user/{user_id}", userHandler.DeleteUser)

	if err := fasthttp.ListenAndServe(c.Host+":"+c.Port, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
