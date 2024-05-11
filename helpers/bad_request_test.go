package helpers_test

import (
	"testing"

	"github.com/proudynyu/clean_arch_go/helpers"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnAStatus400(t *testing.T) {
	e := helpers.BadRequest("Name")

	assert.Equal(t, e.Status, 400)
}

func Test_ShouldReturnMsgPassedAsParam(t *testing.T) {
	e := helpers.BadRequest("Name")

	assert.Equal(t, e.Msg, "Missing Name")
}
