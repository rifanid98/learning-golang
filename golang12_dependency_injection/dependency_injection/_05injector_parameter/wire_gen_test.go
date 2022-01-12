package _045injector_parameter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleServiceSuccess(t *testing.T) {
	service, err := InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, service)
}

func TestSimpleServiceError(t *testing.T) {
	service, err := InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, service)
}
