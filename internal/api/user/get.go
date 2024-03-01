package user

import (
	"basic-backend/internal/converter"
	"basic-backend/internal/entities"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (i *API) Get(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  context.Context
		uuid entities.UUID
	)

	ctx = context.Background()

	values := r.URL.Query()
	uuid = entities.UUID(values.Get("uuid"))

	user, err := i.userService.Get(ctx, uuid)
	if err != nil {
		log.Printf("user: %v", err)
		fmt.Fprintf(w, "%v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	bs, err := json.Marshal(converter.ToUserFromService(&user))
	if err != nil {
		log.Printf("user: %v", err)
		fmt.Fprintf(w, "%v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(bs)
	if err != nil {
		log.Printf("user: %v", err)
		fmt.Fprintf(w, "%v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
}
