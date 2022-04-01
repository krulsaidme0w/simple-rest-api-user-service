package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"golang_pet_project_1/internal/handlers/userhandler"
	"log"
)

const host = "0.0.0.0"
const port = "8080"

func main() {
	userHandler := userhandler.UserHandler{}

	r := router.New()
	r.POST("/user/", userHandler.CreateUser)
	r.GET("/user/get/", userHandler.GetUser)
	r.POST("/user/{user_id}", userHandler.UpdateUser)
	r.DELETE("/user/{user_id}", userHandler.DeleteUser)

	if err := fasthttp.ListenAndServe(host+":"+port, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
