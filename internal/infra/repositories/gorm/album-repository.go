package gormrepositories

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	gormmodels "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm/models"
	"gorm.io/gorm"
)

type GormAlbumRepository struct {
	db *gorm.DB
}

var _ repositories.AlbumRepository = (*GormAlbumRepository)(nil)

func NewGormAlbumRepository(db *gorm.DB) *GormAlbumRepository {
	return &GormAlbumRepository{
		db,
	}
}

func (r GormAlbumRepository) Insert(entity *entities.Album) error {
	model := entityToAlbumModel(entity)

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func entityToAlbumModel(entity *entities.Album) *gormmodels.AlbumModel {
	if entity == nil {
		return nil
	}

	var musics []gormmodels.MusicModel
	for _, music := range entity.Musics {
		musics = append(musics, *entityToMusicModel(music))
	}

	return &gormmodels.AlbumModel{
		ID:            entity.ID,
		Name:          entity.Name,
		ReleaseDate:   entity.ReleaseDate,
		OwnerID:       entity.OwnerID,
		CoverImageURL: entity.CoverImageURL,
		Owner:         *entityToOwnerModel(entity.Owner),
		Musics:        musics,
	}
}
