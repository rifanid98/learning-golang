package test

import (
	"github.com/stretchr/testify/assert"
	"golang11_dependency_injection/dependency_injection"
	"testing"
)

func TestSimpleServiceSuccess(t *testing.T) {
	service, err := dependency_injection.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, service)
}

func TestSimpleServiceError(t *testing.T) {
	service, err := dependency_injection.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, service)
}
