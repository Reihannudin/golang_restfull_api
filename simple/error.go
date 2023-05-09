package simple

import "errors"

type ErrorRepository struct {
	Error bool
}

type ErrorService struct {
	*ErrorRepository
}

func NewErrorRepository(isError bool) *ErrorRepository {
	return &ErrorRepository{
		Error: isError,
	}
}

func NewErrorService(repository *ErrorRepository) (*ErrorService, error) {
	if repository.Error {
		return nil, errors.New("Failed create service")
	} else {
		return &ErrorService{
			ErrorRepository: repository,
		}, nil
	}
}
