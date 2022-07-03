package main

import "golang_pet_project_1/pkg/scripts"

const userCount = 1000

func main() {
	scripts.FillDB(userCount, "db")
}
