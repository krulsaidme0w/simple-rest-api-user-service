package main

import (
	"log"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"golang_pet_project_1/internal/config"
	"golang_pet_project_1/internal/dig"
)

func main() {
	container, err := dig.GetContainer(os.Getenv("STORAGETYPE"))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := container.Invoke(func(c *config.Config, r *router.Router) {
		if err := fasthttp.ListenAndServe(c.Host+":"+c.Port, r.Handler); err != nil {
			log.Fatal(err.Error())
		}
	}); err != nil {
		log.Fatal(err.Error())
	}
}
