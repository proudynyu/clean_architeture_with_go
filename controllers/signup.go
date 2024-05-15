package controllers

import (
	helpers "github.com/proudynyu/clean_arch_go/helpers"
	protocols "github.com/proudynyu/clean_arch_go/protocols"
)

type SignUpController struct {
	emailValidator protocols.EmailValidatorInterface
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
	if !s.emailValidator.IsValid(request.Body.Email) {
		return helpers.BadRequest("Email")
	}
	return protocols.HttpResponse{}
}

func NewSignUpController(emailValidator protocols.EmailValidatorInterface) *SignUpController {
	return &SignUpController{
		emailValidator,
	}
}
