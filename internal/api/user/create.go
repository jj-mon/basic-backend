package user

import (
	"fmt"
	"net/http"
)

func (i *API) Create(w http.ResponseWriter, r *http.Request) {
	buffer := make([]byte, 0, 1024)
	_, err := r.Body.Read(buffer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	fmt.Fprint(w, "user created")
}
