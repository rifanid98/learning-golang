package dependency_injection

import "errors"

/**
Provider

- Untuk melakukan Dependency Injection, kita perlu buat dalam bentuk function constructor
- Dalam Google Wire, function constructor tersebut kita sebut dengan Provider
*/

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	Repository *SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{Repository: repository}, nil
	}
}
