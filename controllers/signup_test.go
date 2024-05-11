package controllers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	controllers "github.com/proudynyu/clean_arch_go/controllers"
	"github.com/proudynyu/clean_arch_go/helpers"
	protocols "github.com/proudynyu/clean_arch_go/protocols"
)

func Test_ShouldReturn400IfNoNameIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := protocols.HttpRequest{
		Body: protocols.HttpBody{
			Name:                 "",
			Email:                "igor@doe.com",
			Password:             "123",
			PasswordConfirmation: "123_confirm",
		},
	}

	response := sut.Handler(httpRequest)

	e := helpers.BadRequest("Name")
	assert.Equal(t, response.Status, e.Status)
	assert.Equal(t, response.Msg, e.Msg)
}

func Test_ShouldReturn400IfNoEmailIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := protocols.HttpRequest{
		Body: protocols.HttpBody{
			Name:                 "igor",
			Email:                "",
			Password:             "123",
			PasswordConfirmation: "123_confirm",
		},
	}

	response := sut.Handler(httpRequest)

	e := helpers.BadRequest("Email")
	assert.Equal(t, response.Status, e.Status)
	assert.Equal(t, response.Msg, e.Msg)
}

func Test_ShouldReturn400IfNoPasswordIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := protocols.HttpRequest{
		Body: protocols.HttpBody{
			Name:                 "igor",
			Email:                "igor@email.com",
			Password:             "",
			PasswordConfirmation: "123_confirm",
		},
	}

	response := sut.Handler(httpRequest)

	e := helpers.BadRequest("Password")
	assert.Equal(t, response.Status, e.Status)
	assert.Equal(t, response.Msg, e.Msg)
}

func Test_ShouldReturn400IfNoPasswordConfirmationIsPassed(t *testing.T) {
	sut := controllers.NewSignUpController()

	httpRequest := protocols.HttpRequest{
		Body: protocols.HttpBody{
			Name:                 "igor",
			Email:                "igor@email.com",
			Password:             "password",
			PasswordConfirmation: "",
		},
	}

	response := sut.Handler(httpRequest)

	e := helpers.BadRequest("PasswordConfirmation")
	assert.Equal(t, response.Status, e.Status)
	assert.Equal(t, response.Msg, e.Msg)
}
