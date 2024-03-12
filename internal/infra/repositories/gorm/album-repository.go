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

func (r GormAlbumRepository) FindBy(where *entities.Album, includes []string) (*entities.Album, error) {
	var model gormmodels.AlbumModel

	query := r.db

	for _, relation := range includes {
		query = query.Preload(relation)
	}

	if err := query.Where(where).First(&model).Error; err != nil {
		return nil, err
	}

	return modelToAlbumEntity(&model), nil
}

func entityToAlbumModel(entity *entities.Album) *gormmodels.AlbumModel {
	if entity == nil {
		return nil
	}

	var musics []*gormmodels.MusicModel
	for _, music := range entity.Musics {
		musics = append(musics, entityToMusicModel(music))
	}

	return &gormmodels.AlbumModel{
		ID:            entity.ID,
		Name:          entity.Name,
		ReleaseDate:   entity.ReleaseDate,
		OwnerID:       entity.OwnerID,
		CoverImageURL: entity.CoverImageURL,
		Owner:         entityToOwnerModel(entity.Owner),
		Musics:        musics,
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

	return &entities.Album{
		ID:            model.ID,
		Name:          model.Name,
		ReleaseDate:   model.ReleaseDate,
		OwnerID:       model.OwnerID,
		CoverImageURL: model.CoverImageURL,
		Owner:         modelToOwnerEntity(model.Owner),
		Musics:        musics,
	}
}
