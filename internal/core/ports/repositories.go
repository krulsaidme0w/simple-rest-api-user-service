package ports

import "os/user"

type UserRepository interface {
	Save(user user.User) (user.User, error)
	Delete(user user.User) error
	Find(searchType string, search string) (user.User, error)
}
