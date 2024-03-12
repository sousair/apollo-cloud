package appusecases

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
)

type GetOwnerAlbumsUsecase struct {
	albumRepository repositories.AlbumRepository
}

var _ usecases.GetOwnerAlbumsUsecase = (*GetOwnerAlbumsUsecase)(nil)

func NewGetOwnerAlbumsUsecase(albumRepository repositories.AlbumRepository) *GetOwnerAlbumsUsecase {
	return &GetOwnerAlbumsUsecase{
		albumRepository,
	}
}

func (uc GetOwnerAlbumsUsecase) GetAlbums(params usecases.GetOwnerAlbumsParams) ([]*entities.Album, error) {
	includes := []string{}

	if params.IncludeMusicData {
		includes = append(includes, "Musics")
	}

	album, err := uc.albumRepository.FindAllBy(&entities.Album{OwnerID: params.OwnerID}, includes)

	if err != nil {
		return nil, err
	}

	return album, nil
}
