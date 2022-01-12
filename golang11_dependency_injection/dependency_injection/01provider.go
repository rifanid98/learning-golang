package dependency_injection

/**
Provider

- Untuk melakukan Dependency Injection, kita perlu buat dalam bentuk function constructor
- Dalam Google Wire, function constructor tersebut kita sebut dengan Provider
*/

type SimpleRepository struct {
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	repository *SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{repository: repository}
}
