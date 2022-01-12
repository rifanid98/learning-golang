package _07provider_set

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(
	fooService *FooService,
	barService *BarService,
) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barService,
	}
}
