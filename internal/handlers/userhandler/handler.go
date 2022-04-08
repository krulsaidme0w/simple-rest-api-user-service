package userhandler

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/internal/core/ports"
	"golang_pet_project_1/internal/validator"
	"golang_pet_project_1/pkg/errors/handler_errors"
	"golang_pet_project_1/pkg/errors/repository_errors"
)

type UserHandler struct {
	userService ports.UserService
}

func NewHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) CreateUser(ctx *fasthttp.RequestCtx) {
	user := &domain.User{}
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		ctx.Error("Invalid input", fasthttp.StatusMethodNotAllowed)
		return
	}

	createdUser, err := h.userService.Create(user)
	if err != nil {
		ctx.Error("Bad request", fasthttp.StatusBadRequest)
		return
	}

	b, err := json.Marshal(createdUser)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(b)
	ctx.SetStatusCode(200)
}

func (h UserHandler) GetUser(ctx *fasthttp.RequestCtx) {
	searchType := string(ctx.QueryArgs().Peek("searchType"))
	search := string(ctx.QueryArgs().Peek("search"))

	user, err := h.userService.Find(searchType, search)
	if err != nil {
		switch err {
		case handler_errors.InvalidSearchType:
			ctx.Error("Invalid searchType supplied", fasthttp.StatusBadRequest)
		default:
			ctx.Error("User not found", fasthttp.StatusNotFound)
		}
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(b)
	ctx.SetStatusCode(200)
}

func (h UserHandler) UpdateUser(ctx *fasthttp.RequestCtx) {
	user := &domain.User{}
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		ctx.Error("Invalid input", fasthttp.StatusMethodNotAllowed)
		return
	}

	if validateUserAndID := validator.ValidateUserAndPath(user, ctx.UserValue("user_id").(string)); !validateUserAndID {
		ctx.Error("Invalid user id supplied", fasthttp.StatusBadRequest)
		return
	}

	updatedUser, err := h.userService.Update(user)
	if err != nil {
		switch err {
		case repository_errors.UserNotExists:
			ctx.Error("User not found", fasthttp.StatusNotFound)
			return
		}
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(updatedUser)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(b)
	ctx.SetStatusCode(200)
}

func (h UserHandler) DeleteUser(ctx *fasthttp.RequestCtx) {
	err := h.userService.Delete(ctx.UserValue("user_id").(string))
	if err != nil {
		switch err {
		case repository_errors.UserNotExists:
			ctx.Error("User not found", fasthttp.StatusNotFound)
			return
		}
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(200)
	ctx.SetBody([]byte("User deleted"))
}
