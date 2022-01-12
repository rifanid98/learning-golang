package _02injector

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	Repository *SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{Repository: repository}
}
