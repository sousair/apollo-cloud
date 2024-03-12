package usecases

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	GetOwnerAlbumsUsecase interface {
		GetAlbums(params GetOwnerAlbumsParams) ([]*entities.Album, error)
	}

	GetOwnerAlbumsParams struct {
		OwnerID          string
		IncludeMusicData bool
	}
)
