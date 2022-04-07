package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bxcodec/faker/v3"

	"golang_pet_project_1/internal/core/domain"
)

func FillDB(userCount int) {
	maxGoroutines := 10
	guard := make(chan struct{}, maxGoroutines)

	for i := 0; i < userCount; i++ {
		guard <- struct{}{}
		go func(i int) {
			start := time.Now()

			var user domain.User
			err := faker.FakeData(&user)
			user.ID = i
			if err != nil {
				fmt.Println(err)
			}

			b, err := json.Marshal(user)
			if err != nil {
				fmt.Println(err)
			}

			resp, err := http.Post("http://0.0.0.0:8080/userservice/", "application/json", bytes.NewReader(b))
			if err != nil {
				fmt.Println(err)
			}

			err = resp.Body.Close()
			if err != nil {
				fmt.Println(err)
			}

			stop := time.Now()

			fmt.Print("TIME: ")
			fmt.Println(stop.Sub(start).Milliseconds())

			<-guard
		}(i)
	}
}
