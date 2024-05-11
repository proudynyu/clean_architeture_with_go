package controllers_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	controllers "github.com/proudynyu/clean_arch_go/controllers"
	utils "github.com/proudynyu/clean_arch_go/utils"
)

func Test_ShouldReturn400IfNoNameIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := utils.HttpRequest{
		Body: utils.HttpBody{
			Name:                 "",
			Email:                "igor@doe.com",
			Password:             "123",
			PasswordConfirmation: "123_confirm",
		},
	}

	response := sut.Handler(httpRequest)

	assert.Equal(t, response.Status, 400)
	assert.Equal(t, response.Msg, utils.ErrorName)
}

func Test_ShouldReturn400IfNoEmailIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := utils.HttpRequest{
		Body: utils.HttpBody{
			Name:                 "igor",
			Email:                "",
			Password:             "123",
			PasswordConfirmation: "123_confirm",
		},
	}

	response := sut.Handler(httpRequest)

	assert.Equal(t, response.Status, 400)
	assert.Equal(t, response.Msg, utils.ErrorEmail)
}

func Test_ShouldReturn400IfNoPasswordIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := utils.HttpRequest{
		Body: utils.HttpBody{
			Name:                 "igor",
			Email:                "igor@email.com",
			Password:             "",
			PasswordConfirmation: "123_confirm",
		},
	}

	response := sut.Handler(httpRequest)

	assert.Equal(t, response.Status, 400)
	assert.Equal(t, response.Msg, utils.ErrorPassword)
}

func Test_ShouldReturn400IfNoPasswordConfirmationIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := utils.HttpRequest{
		Body: utils.HttpBody{
			Name:                 "igor",
			Email:                "igor@email.com",
			Password:             "password",
			PasswordConfirmation: "",
		},
	}

	response := sut.Handler(httpRequest)

	assert.Equal(t, response.Status, 400)
	assert.Equal(t, response.Msg, utils.ErrorPasswordConfirmation)
}
