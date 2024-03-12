package appusecases

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/providers"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
)

type CreateOwnerUsecase struct {
	uuidProvider    providers.UuidProvider
	ownerRepository repositories.OwnerRepository
}

var _ usecases.CreateOwnerUsecase = (*CreateOwnerUsecase)(nil)

func NewCreateOwnerUsecase(
	uuidProvider providers.UuidProvider,
	ownerRepository repositories.OwnerRepository,
) *CreateOwnerUsecase {
	return &CreateOwnerUsecase{
		uuidProvider,
		ownerRepository,
	}
}

func (uc CreateOwnerUsecase) Create(params usecases.CreateOwnerParams) (*entities.Owner, error) {
	owner, err := entities.NewOwner(entities.NewOwnerParams{
		ID:   uc.uuidProvider.Generate(),
		Name: params.Name,
	})

	if err != nil {
		return nil, err
	}

	if err := uc.ownerRepository.Insert(owner); err != nil {
		return nil, err
	}

	return owner, nil
}
