package ports

import "os/user"

type UserService interface {
	Create(user user.User) (user.User, error)
	Find(searchType string, search string) (user.User, error)
	Update(user user.User) (user.User, error)
	Delete(user user.User) error
}
