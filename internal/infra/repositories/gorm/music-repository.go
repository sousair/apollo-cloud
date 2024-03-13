package gormrepositories

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/valueobjects"
	gormmodels "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm/models"
	"gorm.io/gorm"
)

type GormMusicRepository struct {
	db *gorm.DB
}

var _ repositories.MusicRepository = (*GormMusicRepository)(nil)

func NewGormMusicRepository(db *gorm.DB) *GormMusicRepository {
	db = db.Preload("CoverImageLocation").Preload("MusicFileLocation")
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

func (r GormMusicRepository) FindBy(where map[string]interface{}, includes []string) (*entities.Music, error) {
	var model gormmodels.MusicModel

	query := r.db

	for _, relation := range includes {
		query = query.Preload(relation)

		// gorm is bad.
		if relation == "Album" {
			query = query.Preload("Album.CoverImageLocation")
		}
	}

	if err := query.Where(where).First(&model).Error; err != nil {
		return nil, err
	}

	return modelToMusicEntity(&model), nil
}

func entityToMusicModel(entity *entities.Music) *gormmodels.MusicModel {
	if entity == nil {
		return nil
	}

	var coverImageLocation *gormmodels.FileLocationModel
	if entity.CoverImageLocation != nil {
		coverImageLocation = &gormmodels.FileLocationModel{
			URL:       entity.CoverImageLocation.URL,
			Provider:  entity.CoverImageLocation.Provider,
			Extension: entity.CoverImageLocation.Extension,
		}
	}

	musicFileLocation := &gormmodels.FileLocationModel{
		URL:       entity.MusicFileLocation.URL,
		Provider:  entity.MusicFileLocation.Provider,
		Extension: entity.MusicFileLocation.Extension,
	}

	return &gormmodels.MusicModel{
		ID:                 entity.ID,
		Name:               entity.Name,
		DurationInMs:       entity.DurationInMs,
		ReleaseDate:        entity.ReleaseDate,
		AlbumID:            entity.AlbumID,
		OwnerID:            entity.OwnerID,
		CoverImageLocation: coverImageLocation,
		MusicFileLocation:  musicFileLocation,
		Album:              entityToAlbumModel(entity.Album),
		Owner:              entityToOwnerModel(entity.Owner),
	}
}

func modelToMusicEntity(model *gormmodels.MusicModel) *entities.Music {
	if model == nil {
		return nil
	}

	var coverImageLocation *valueobjects.FileLocation
	if model.CoverImageLocation != nil {
		coverImageLocation = &valueobjects.FileLocation{
			URL:       model.CoverImageLocation.URL,
			Provider:  model.CoverImageLocation.Provider,
			Extension: model.CoverImageLocation.Extension,
		}
	}

	musicFileLocation := &valueobjects.FileLocation{
		URL:       model.MusicFileLocation.URL,
		Provider:  model.MusicFileLocation.Provider,
		Extension: model.MusicFileLocation.Extension,
	}

	return &entities.Music{
		ID:                 model.ID,
		Name:               model.Name,
		DurationInMs:       model.DurationInMs,
		ReleaseDate:        model.ReleaseDate,
		AlbumID:            model.AlbumID,
		OwnerID:            model.OwnerID,
		CoverImageLocation: coverImageLocation,
		MusicFileLocation:  musicFileLocation,
		Album:              modelToAlbumEntity(model.Album),
		Owner:              modelToOwnerEntity(model.Owner),
	}
}
