package usecases

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	CreateOwnerUsecase interface {
		Create(params CreateOwnerParams) (*entities.Owner, error)
	}

	CreateOwnerParams struct {
		Name string
	}
)
