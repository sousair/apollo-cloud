package gormrepositories

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	gormmodels "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm/models"
	"gorm.io/gorm"
)

type GormOwnerRepository struct {
	db *gorm.DB
}

var _ repositories.OwnerRepository = (*GormOwnerRepository)(nil)

func NewGormOwnerRepository(db *gorm.DB) *GormOwnerRepository {
	return &GormOwnerRepository{
		db,
	}
}

func (r GormOwnerRepository) Insert(entity *entities.Owner) error {
	model := entityToOwnerModel(entity)

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func entityToOwnerModel(entity *entities.Owner) *gormmodels.OwnerModel {
	if entity == nil {
		return nil
	}

	var albums []*gormmodels.AlbumModel
	for _, album := range entity.Albums {
		albums = append(albums, entityToAlbumModel(album))
	}

	var musics []*gormmodels.MusicModel
	for _, music := range entity.Musics {
		musics = append(musics, entityToMusicModel(music))
	}

	return &gormmodels.OwnerModel{
		ID:     entity.ID,
		Name:   entity.Name,
		Albums: albums,
		Musics: musics,
	}
}

func modelToOwnerEntity(model *gormmodels.OwnerModel) *entities.Owner {
	if model == nil {
		return nil
	}

	var albums []*entities.Album
	for _, album := range model.Albums {
		albums = append(albums, modelToAlbumEntity(album))
	}

	var musics []*entities.Music
	for _, music := range model.Musics {
		musics = append(musics, modelToMusicEntity(music))
	}

	return &entities.Owner{
		ID:     model.ID,
		Name:   model.Name,
		Albums: albums,
		Musics: musics,
	}
}
