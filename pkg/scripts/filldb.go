package scripts

import (
	"log"
	"strconv"
	"sync"

	"github.com/bxcodec/faker/v3"

	"golang_pet_project_1/internal/core/domain"
)

func FillDB(userCount int, path string) {
	guardChan := make(chan struct{}, 2)
	wg := &sync.WaitGroup{}
	for i := 1; i < userCount+1; i++ {
		guardChan <- struct{}{}
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			var user domain.User
			err := faker.FakeData(&user)
			user.ID = i
			if err != nil {
				log.Fatalf(err.Error())
			}

			_, err = domain.WriteUserToFile(path+"/"+strconv.Itoa(user.ID), &user)
			if err != nil {
				log.Fatal(err.Error())
			}
			<-guardChan
		}(i, wg)
	}
	wg.Done()
}
