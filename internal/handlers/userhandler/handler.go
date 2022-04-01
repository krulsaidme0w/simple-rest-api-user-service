package userhandler

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"golang_pet_project_1/internal/core/domain"
)

type UserHandler struct {
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

}

func (UserHandler) UpdateUser(ctx *fasthttp.RequestCtx) {

}

func (UserHandler) DeleteUser(ctx *fasthttp.RequestCtx) {

}
