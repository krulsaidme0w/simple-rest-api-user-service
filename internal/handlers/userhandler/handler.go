package userhandler

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/internal/core/ports"
)

type UserHandler struct {
	userService ports.UserService
}

func (UserHandler) CreateUser(ctx *fasthttp.RequestCtx) {
	user := &domain.User{}
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		ctx.Error("Invalid input", fasthttp.StatusMethodNotAllowed)
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

func (UserHandler) GetUser(ctx *fasthttp.RequestCtx) {

	searchType := string(ctx.QueryArgs().Peek("searchType"))
	//search := string(ctx.QueryArgs().Peek("search"))

	user := &domain.User{
		ID:       1,
		Username: "asd",
		Name:     "clown",
		Age:      123,
	}

	b, err := json.Marshal(user)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	switch searchType {
	case "id":
		break
	case "username":
		break
	case "name":
		break
	default:
		ctx.Error("Invalid searchType supplied", fasthttp.StatusBadRequest)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(b)
}

func (UserHandler) UpdateUser(ctx *fasthttp.RequestCtx) {

}

func (UserHandler) DeleteUser(ctx *fasthttp.RequestCtx) {

}
