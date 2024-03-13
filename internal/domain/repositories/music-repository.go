package repositories

import "github.com/sousair/apollo-cloud/internal/domain/entities"

type (
	MusicRepository interface {
		Insert(*entities.Music) error
		FindBy(where map[string]interface{}, includes []string) (*entities.Music, error)
	}
)
