package controllers

import "github.com/proudynyu/clean_arch_go/utils"

type SignUpController struct {
}

func (s *SignUpController) Handler(request utils.HttpRequest) utils.HttpResponse {
	if request.Body.Name == "" {
		return utils.HttpResponse{
			Status: 400,
			Msg:    utils.ErrorName,
		}
	}
	if request.Body.Email == "" {
		return utils.HttpResponse{
			Status: 400,
			Msg:    utils.ErrorEmail,
		}
	}
	if request.Body.Password == "" {
		return utils.HttpResponse{
			Status: 400,
			Msg:    utils.ErrorPassword,
		}
	}
	if request.Body.PasswordConfirmation == "" {
		return utils.HttpResponse{
			Status: 400,
			Msg:    utils.ErrorPasswordConfirmation,
		}
	}

	return utils.HttpResponse{}
}

func NewSignUpController() *SignUpController {
	return &SignUpController{}
}
