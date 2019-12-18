package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type UserResource struct {
	logger  *zap.Logger
	service *userService
}

func NewUserResource(l *zap.Logger) UserResource {
	us := newUserService()
	ur := UserResource{logger: l, service: us}
	return ur
}

func (u UserResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", u.registerUser())
	return r
}

func (u UserResource) registerUser() http.HandlerFunc {

	type userRequest struct {
		User User `json:"user"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user userRequest
		// b := r.Body

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		regUser, err := u.service.RegisterUser(user.User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		js, err := json.Marshal(regUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

}
