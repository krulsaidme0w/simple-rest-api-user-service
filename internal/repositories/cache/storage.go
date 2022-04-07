package cache

import (
	"io/ioutil"

	iradix "github.com/hashicorp/go-immutable-radix"

	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/pkg/errors/repository_errors"
)

type Cache struct {
	usernamePrefix *iradix.Tree
	namePrefix     *iradix.Tree
}

func NewCache(usernamePrefix *iradix.Tree, namePrefix *iradix.Tree) *Cache {
	return &Cache{
		usernamePrefix: usernamePrefix,
		namePrefix:     namePrefix,
	}
}

func (c *Cache) FillFromDB(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		user, err := domain.ReadUserFromFile(path + "/" + f.Name())
		if err != nil {
			return err
		}

		c.usernamePrefix, _, _ = c.usernamePrefix.Insert([]byte(user.Username), user.ID)
		c.namePrefix, _, _ = c.namePrefix.Insert([]byte(user.Name), user.ID)
	}

	return nil
}

func (c *Cache) Add(user *domain.User) error {
	c.usernamePrefix, _, _ = c.usernamePrefix.Insert([]byte(user.Username), user.ID)
	c.namePrefix, _, _ = c.namePrefix.Insert([]byte(user.Name), user.ID)

	return nil
}

func (c *Cache) GetIDByUsername(username string) (int, error) {
	id, isFound := c.usernamePrefix.Get([]byte(username))
	if !isFound {
		return 0, repository_errors.CannotFindUserInCache
	}
	return id.(int), nil
}

func (c *Cache) GetIDByName(name string) (int, error) {
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
	c.usernamePrefix, _, _ = c.usernamePrefix.Delete([]byte(username))

	return nil
}
