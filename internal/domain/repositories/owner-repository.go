package repositories

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	OwnerRepository interface {
		Insert(*entities.Owner) error
	}
)
