package helpers

import "github.com/proudynyu/clean_arch_go/protocols"

func BadRequest(msg string) protocols.HttpResponse {
	return protocols.HttpResponse{
		Status: 400,
		Msg:    "Missing " + msg,
	}
}
