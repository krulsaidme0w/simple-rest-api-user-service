package repository_errors

import "errors"

var CannotFindUser = errors.New("CannotFindUser")
var CannotAddIndex = errors.New("CannotAddIndex")
var UserAlreadyExists = errors.New("UserAlreadyExists")
var UserNotExists = errors.New("UserNotExists")

var CannotAddOrUpdateUserToCache = errors.New("CannotAddUserToCache")
var CannotFindUserInCache = errors.New("CannotFindUserInCache")
var CannotDeleteUserFromCache = errors.New("CannotDeleteUserFromCache")

var UserWithThisUsernameExists = errors.New("UserWithThisUsernameExists")
var UserWithThisNameExists = errors.New("UserWithThisNameExists")
