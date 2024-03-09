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

func (uc CreateMusicUsecase) Create(params usecases.CreateMusicParams) (*entities.Music, error) {
	coverImageLocation, err := uc.fileRepository.Upload(repositories.UploadFileParams{
		File:   params.CoverImage,
		Public: true,
	})

	if err != nil {
		return nil, err
	}

	songLocation, err := uc.fileRepository.Upload(repositories.UploadFileParams{
		File:   params.Song,
		Public: false,
	})

	if err != nil {
		return nil, err
	}

	music, err := entities.NewMusic(entities.NewMusicParams{
		ID:               uc.uuidProvider.Generate(),
		Name:             params.Name,
		OwnerID:          params.OwnerID,
		DurationInMs:     params.DurationInMs,
		ReleaseDate:      params.ReleaseDate,
		CoverImageURL:    coverImageLocation.URL,
		SongDataLocation: songLocation.URL,
	})

	if params.AlbumID != "" {
		music.AlbumID = params.AlbumID
	}

	if err = uc.musicRepository.Insert(music); err != nil {
		return nil, err
	}

	return music, nil
}
