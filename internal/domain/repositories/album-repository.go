package repositories

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	AlbumRepository interface {
		Insert(*entities.Album) error
		FindBy(where *entities.Album, includes []string) (*entities.Album, error)
	}
)
