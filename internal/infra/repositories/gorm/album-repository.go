package gormrepositories

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/valueobjects"
	gormmodels "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm/models"
	"gorm.io/gorm"
)

type GormAlbumRepository struct {
	db *gorm.DB
}

var _ repositories.AlbumRepository = (*GormAlbumRepository)(nil)

func NewGormAlbumRepository(db *gorm.DB) *GormAlbumRepository {
	db = db.Preload("CoverImageLocation")
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

func (r GormAlbumRepository) FindBy(where map[string]interface{}, includes []string) (*entities.Album, error) {
	var model gormmodels.AlbumModel

	query := r.db

	for _, relation := range includes {
		query = query.Preload(relation)

		// gorm is bad.
		if relation == "Musics" {
			query = query.Preload("Musics.CoverImageLocation").Preload("Musics.MusicFileLocation")
		}
	}

	if err := query.Where(where).First(&model).Error; err != nil {
		return nil, err
	}

	return modelToAlbumEntity(&model), nil

}

func (r GormAlbumRepository) FindAllBy(where map[string]interface{}, includes []string) ([]*entities.Album, error) {
	var models []*gormmodels.AlbumModel

	query := r.db

	for _, relation := range includes {
		query = query.Preload(relation)

		// gorm is bad.
		if relation == "Musics" {
			query = query.Preload("Musics.CoverImageLocation").Preload("Musics.MusicFileLocation")
		}
	}

	if err := query.Where(where).Find(&models).Error; err != nil {
		return nil, err
	}

	var entities []*entities.Album
	for _, model := range models {
		entities = append(entities, modelToAlbumEntity(model))
	}

	return entities, nil
}

func entityToAlbumModel(entity *entities.Album) *gormmodels.AlbumModel {
	if entity == nil {
		return nil
	}

	var musics []*gormmodels.MusicModel
	for _, music := range entity.Musics {
		musics = append(musics, entityToMusicModel(music))
	}

	coverLocation := &gormmodels.FileLocationModel{
		URL:       entity.CoverImageLocation.URL,
		Provider:  entity.CoverImageLocation.Provider,
		Extension: entity.CoverImageLocation.Extension,
	}

	return &gormmodels.AlbumModel{
		ID:                 entity.ID,
		Name:               entity.Name,
		ReleaseDate:        entity.ReleaseDate,
		OwnerID:            entity.OwnerID,
		CoverImageLocation: coverLocation,
		Owner:              entityToOwnerModel(entity.Owner),
		Musics:             musics,
	}
}

func modelToAlbumEntity(model *gormmodels.AlbumModel) *entities.Album {
	if model == nil {
		return nil
	}

	var musics []*entities.Music
	for _, music := range model.Musics {
		musics = append(musics, modelToMusicEntity(music))
	}

	coverLocation := &valueobjects.FileLocation{
		URL:       model.CoverImageLocation.URL,
		Provider:  model.CoverImageLocation.Provider,
		Extension: model.CoverImageLocation.Extension,
	}

	return &entities.Album{
		ID:                 model.ID,
		Name:               model.Name,
		ReleaseDate:        model.ReleaseDate,
		OwnerID:            model.OwnerID,
		CoverImageLocation: coverLocation,
		Owner:              modelToOwnerEntity(model.Owner),
		Musics:             musics,
	}
}
