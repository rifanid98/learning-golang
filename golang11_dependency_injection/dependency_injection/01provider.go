package dependency_injection

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
