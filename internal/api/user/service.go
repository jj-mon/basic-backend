package user

import (
	"basic-backend/internal/service"
	"log"
	"net/http"
)

type API struct {
	handlers    map[string]http.HandlerFunc
	userService service.UserService
}

func New(userService service.UserService) *API {
	i := &API{
		userService: userService,
	}

	i.initHandlers()

	return i
}

func (i *API) Handlers() map[string]http.HandlerFunc {
	return i.handlers
}

func (i *API) initHandlers() {
	i.handlers = make(map[string]http.HandlerFunc)
	i.handlers["/user"] = i.Get
	i.handlers["POST /user"] = i.Create

	log.Println("init user api")
}
