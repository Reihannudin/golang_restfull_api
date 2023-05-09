package simple

// Foo
type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{
		FooRepository: fooRepository,
	}
}

// Bar
type BarRespository struct {
}

func NewBarRepository() *BarRespository {
	return &BarRespository{}
}

type BarService struct {
	*BarRespository
}

func NewBarService(barRespository *BarRespository) *BarService {
	return &BarService{
		BarRespository: barRespository,
	}
}

//Provider set FooBar

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barService,
	}
}
