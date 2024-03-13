package appusecases

import (
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/providers"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
)

type CreateMusicUsecase struct {
	uuidProvider    providers.UuidProvider
	fileRepository  repositories.FileRepository
	musicRepository repositories.MusicRepository
}

var _ usecases.CreateMusicUsecase = (*CreateMusicUsecase)(nil)

func NewCreateMusicUsecase(
	uuidProvider providers.UuidProvider,
	fileRepository repositories.FileRepository,
	musicRepository repositories.MusicRepository,
) *CreateMusicUsecase {
	return &CreateMusicUsecase{
		uuidProvider,
		fileRepository,
		musicRepository,
	}
}

func (uc CreateMusicUsecase) Create(params usecases.CreateMusicParams) (music *entities.Music, err error) {
	musicFileLocation, err := uc.fileRepository.Upload(repositories.UploadFileParams{
		File:   params.MusicFile,
		Public: false,
	})

	if err != nil {
		return nil, err
	}

	music, err = entities.NewMusic(entities.NewMusicParams{
		ID:                uc.uuidProvider.Generate(),
		Name:              params.Name,
		AlbumID:           params.AlbumID,
		OwnerID:           params.OwnerID,
		DurationInMs:      params.DurationInMs,
		ReleaseDate:       params.ReleaseDate,
		MusicFileLocation: musicFileLocation,
	})

	if err != nil {
		return nil, err
	}

	if params.CoverImage != nil {
		coverImageLocation, err := uc.fileRepository.Upload(repositories.UploadFileParams{
			File:   params.CoverImage,
			Public: true,
		})

		if err != nil {
			return nil, err
		}

		music.CoverImageLocation = coverImageLocation
	}

	if err = uc.musicRepository.Insert(music); err != nil {
		return nil, err
	}

	return music, nil
}
