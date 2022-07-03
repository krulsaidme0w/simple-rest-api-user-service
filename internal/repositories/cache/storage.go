package cache

import (
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

func NewCache() *Cache {
	mutex := &sync.RWMutex{}
	return &Cache{
		usernamePrefix: iradix.New(),
		namePrefix:     iradix.New(),
		mutex:          mutex,
	}
}

func (c *Cache) Fill(users []domain.User) {
	for _, user := range users {
		c.usernamePrefix, _, _ = c.usernamePrefix.Insert([]byte(user.Username), user.ID)
		c.namePrefix, _, _ = c.namePrefix.Insert([]byte(user.Name), user.ID)
	}
}

func (c *Cache) Add(user *domain.User) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, isFound := c.usernamePrefix.Get([]byte(user.Username))
	if !isFound {
		return repository_errors.UserWithThisUsernameExists
	}

	_, isFound = c.namePrefix.Get([]byte(user.Name))
	if !isFound {
		return repository_errors.UserWithThisNameExists
	}

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
