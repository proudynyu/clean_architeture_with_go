package protocols

type HttpBody struct {
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

type HttpRequest struct {
	Body HttpBody
}

type HttpResponse struct {
	Status int
	Msg    string
}
