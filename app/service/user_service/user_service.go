package user_service

import "fiber-scaffold/app/repository"

type Auth struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

func (auth *Auth) Check() bool {
	return repository.CheckLoginUser(auth.Username, auth.Password)
}
