package controllers

import (
	helpers "github.com/proudynyu/clean_arch_go/helpers"
	protocols "github.com/proudynyu/clean_arch_go/protocols"
)

type SignUpController struct {
}

func (s *SignUpController) Handler(request protocols.HttpRequest) protocols.HttpResponse {
	if request.Body.Name == "" {
		return helpers.BadRequest("Name")
	}
	if request.Body.Email == "" {
		return helpers.BadRequest("Email")
	}
	if request.Body.Password == "" {
		return helpers.BadRequest("Password")
	}
	if request.Body.PasswordConfirmation == "" {
		return helpers.BadRequest("PasswordConfirmation")
	}
	return protocols.HttpResponse{}
}

func NewSignUpController() *SignUpController {
	return &SignUpController{}
}
