package cache

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	iradix "github.com/hashicorp/go-immutable-radix"

	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/pkg/errors/repository_errors"
)

type Cache struct {
	usernamePrefix *iradix.Tree
	namePrefix     *iradix.Tree
	mutex          *sync.RWMutex
}

func NewCache(usernamePrefix *iradix.Tree, namePrefix *iradix.Tree, mutex *sync.RWMutex) *Cache {
	return &Cache{
		usernamePrefix: usernamePrefix,
		namePrefix:     namePrefix,
		mutex:          mutex,
	}
}

func (c *Cache) FillFromDB(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	maxGoroutines := 1000
	guardChan := make(chan struct{}, maxGoroutines)

	for _, f := range files {
		wg.Add(1)
		f := f
		guardChan <- struct{}{}

		go func() {
			defer wg.Done()
			user, err := domain.ReadUserFromFile(path + "/" + f.Name())
			if err != nil {
				log.Fatal("bad path:", path+"/"+f.Name(), " ", err)
			}

			c.mutex.Lock()
			c.usernamePrefix, _, _ = c.usernamePrefix.Insert([]byte(user.Username), user.ID)
			c.namePrefix, _, _ = c.namePrefix.Insert([]byte(user.Name), user.ID)
			c.mutex.Unlock()

			<-guardChan
		}()
	}
	close(guardChan)
	wg.Wait()

	fmt.Println("username prefix length: ", c.usernamePrefix.Len())
	fmt.Println("name prefix length: ", c.namePrefix.Len())

	return nil
}

func (c *Cache) Add(user *domain.User) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.usernamePrefix, _, _ = c.usernamePrefix.Insert([]byte(user.Username), user.ID)
	c.namePrefix, _, _ = c.namePrefix.Insert([]byte(user.Name), user.ID)

	return nil
}

func (c *Cache) GetIDByUsername(username string) (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	id, isFound := c.usernamePrefix.Get([]byte(username))
	if !isFound {
		return 0, repository_errors.CannotFindUserInCache
	}
	return id.(int), nil
}

func (c *Cache) GetIDByName(name string) (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	id, isFound := c.namePrefix.Get([]byte(name))
	if !isFound {
		return 0, repository_errors.CannotFindUserInCache
	}
	return id.(int), nil
}

func (c *Cache) Update(user *domain.User) error {
	return c.Add(user)
}

func (c *Cache) Delete(username string, name string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.usernamePrefix, _, _ = c.usernamePrefix.Delete([]byte(username))
	c.namePrefix, _, _ = c.namePrefix.Delete([]byte(name))

	return nil
}
