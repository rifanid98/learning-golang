package _07provider_set

type FooRepository struct{}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	Repository *FooRepository
}

func NewFooService(repository *FooRepository) *FooService {
	return &FooService{Repository: repository}
}
