package simple

// Struct
type SimpleRepository struct {
}
type SimpleService struct {
	*SimpleRepository
}

// Provider

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
