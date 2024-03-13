package repositories

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	AlbumRepository interface {
		Insert(*entities.Album) error
		FindBy(where map[string]interface{}, includes []string) (*entities.Album, error)
		FindAllBy(where map[string]interface{}, includes []string) ([]*entities.Album, error)
	}
)
