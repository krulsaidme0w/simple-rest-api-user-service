package userhandler

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/internal/core/ports"
	"golang_pet_project_1/pkg/errors/handler_errors"
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
}

func (h UserHandler) UpdateUser(ctx *fasthttp.RequestCtx) {

}

func (h UserHandler) DeleteUser(ctx *fasthttp.RequestCtx) {

}
