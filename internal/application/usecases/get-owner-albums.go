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

	albums, err := uc.albumRepository.FindAllBy(
		map[string]interface{}{
			"owner_id": params.OwnerID,
		},
		includes,
	)

	if err != nil {
		return nil, err
	}

	return albums, nil
}
