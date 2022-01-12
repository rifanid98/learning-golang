package test

import (
	"fmt"
	"golang11_dependency_injection/dependency_injection"
	"testing"
)

func TestSimpleService(t *testing.T) {
	service, err := dependency_injection.InitializedService()
	if err != nil {
		return
	}
	fmt.Println(service.Repository)
}
