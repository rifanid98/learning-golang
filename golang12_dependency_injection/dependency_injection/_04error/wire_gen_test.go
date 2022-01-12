package _04error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	service, err := InitializedService()
	assert.Nil(t, err)
	assert.NotNil(t, service)
}
