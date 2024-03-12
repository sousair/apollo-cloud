package appusecases

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
)

type (
	GetAlbumUsecase struct {
		albumRepository repositories.AlbumRepository
	}
)

var _ usecases.GetAlbumUsecase = (*GetAlbumUsecase)(nil)

func NewGetAlbumUsecase(albumRepository repositories.AlbumRepository) *GetAlbumUsecase {
	return &GetAlbumUsecase{
		albumRepository: albumRepository,
	}
}

func (u GetAlbumUsecase) Get(params usecases.GetAlbumParams) (*entities.Album, error) {
	includes := []string{}

	if params.IncludeMusicsData {
		includes = append(includes, "Musics")
	}

	if params.IncludeOwnerData {
		includes = append(includes, "Owner")
	}

	album, err := u.albumRepository.FindBy(&entities.Album{ID: params.ID}, includes)

	if err != nil {
		return nil, err
	}

	return album, nil
}