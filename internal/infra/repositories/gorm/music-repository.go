package gormrepositories

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	gormmodels "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm/models"
	"gorm.io/gorm"
)

type GormMusicRepository struct {
	db *gorm.DB
}

var _ repositories.MusicRepository = (*GormMusicRepository)(nil)

func NewGormMusicRepository(db *gorm.DB) *GormMusicRepository {
	return &GormMusicRepository{
		db,
	}
}

func (r GormMusicRepository) Insert(entity *entities.Music) error {
	model := entityToMusicModel(entity)

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func entityToMusicModel(entity *entities.Music) *gormmodels.MusicModel {
	if entity == nil {
		return nil
	}

	return &gormmodels.MusicModel{
		ID:               entity.ID,
		Name:             entity.Name,
		DurationInMs:     entity.DurationInMs,
		ReleaseDate:      entity.ReleaseDate,
		AlbumID:          entity.AlbumID,
		OwnerID:          entity.OwnerID,
		CoverImageURL:    entity.CoverImageURL,
		SongDataLocation: entity.SongDataLocation,
		Album:            entityToAlbumModel(entity.Album),
		Owner:            entityToOwnerModel(entity.Owner),
	}
}

func modelToMusicEntity(model *gormmodels.MusicModel) *entities.Music {
	if model == nil {
		return nil
	}

	return &entities.Music{
		ID:               model.ID,
		Name:             model.Name,
		DurationInMs:     model.DurationInMs,
		ReleaseDate:      model.ReleaseDate,
		AlbumID:          model.AlbumID,
		OwnerID:          model.OwnerID,
		CoverImageURL:    model.CoverImageURL,
		SongDataLocation: model.SongDataLocation,
		Album:            modelToAlbumEntity(model.Album),
		Owner:            modelToOwnerEntity(model.Owner),
	}
}
