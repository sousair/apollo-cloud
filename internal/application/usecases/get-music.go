package appusecases

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
)

type (
	GetMusicUsecase struct {
		musicRepository repositories.MusicRepository
	}
)

var _ usecases.GetMusicUsecase = (*GetMusicUsecase)(nil)

func NewGetMusicUsecase(musicRepository repositories.MusicRepository) *GetMusicUsecase {
	return &GetMusicUsecase{
		musicRepository,
	}
}

func (uc GetMusicUsecase) Get(params usecases.GetMusicParams) (*entities.Music, error) {
	includes := []string{}

	if params.IncludeAlbumData {
		includes = append(includes, "Album")
	}

	if params.IncludeOwnerData {
		includes = append(includes, "Owner")
	}

	music, err := uc.musicRepository.FindBy(&entities.Music{ID: params.ID}, includes)

	if err != nil {
		return nil, err
	}

	return music, nil
}
