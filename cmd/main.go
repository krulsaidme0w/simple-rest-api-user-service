package main

import (
	"log"
	"sync"

	"github.com/fasthttp/router"
	iradix "github.com/hashicorp/go-immutable-radix"
	"github.com/valyala/fasthttp"

	"golang_pet_project_1/internal/config"
	"golang_pet_project_1/internal/core/services/userservice"
	"golang_pet_project_1/internal/handlers/userhandler"
	userrepository "golang_pet_project_1/internal/repositories"
	"golang_pet_project_1/internal/repositories/cache"
	"golang_pet_project_1/internal/repositories/file"
)

func main() {
	c, err := config.SetUp()
	if err != nil {
		log.Fatal(err.Error())
	}

	//c.DB = "db"
	//c.Port = "8080"
	//c.Host = "0.0.0.0"
	//c.MinioRootUser = "minio"
	//c.MinioRootPassword = "minio123"
	//c.UserBucketName = "users"

	//fileMutex := &sync.RWMutex{}
	prefixMutex := &sync.RWMutex{}

	usernamePrefix := iradix.New()
	namePrefix := iradix.New()

	userCache := cache.NewCache(usernamePrefix, namePrefix, prefixMutex)
	//start := time.Now()
	//err = userCache.FillFromDB(c.DB)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//end := time.Now()
	//
	//fmt.Println("FILLED FROM DB: ", end.Sub(start).Seconds(), "s")

	//fsStorage := file.NewStorage(c.DB, fileMutex)
	minioStorage, err := file.NewMinioStorage(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := userrepository.NewStorage(minioStorage, userCache)
	userService := userservice.NewService(userRepository)
	userHandler := userhandler.NewHandler(userService)

	r := router.New()
	r.POST("/user/", userHandler.CreateUser)
	r.GET("/user/get/", userHandler.GetUser)
	r.POST("/user/{user_id}/", userHandler.UpdateUser)
	r.DELETE("/user/{user_id}/", userHandler.DeleteUser)

	if err := fasthttp.ListenAndServe(c.Host+":"+c.Port, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
