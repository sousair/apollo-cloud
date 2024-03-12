package appusecases

import (
	"fmt"
	"time"

	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/providers"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
)

type ReleaseAlbumUsecase struct {
	fileRepository     repositories.FileRepository
	uuidProvider       providers.UuidProvider
	albumRepository    repositories.AlbumRepository
	createMusicUsecase usecases.CreateMusicUsecase
}

var _ usecases.ReleaseAlbumUsecase = (*ReleaseAlbumUsecase)(nil)

func NewReleaseAlbumUsecase(
	fileRepository repositories.FileRepository,
	uuidProvider providers.UuidProvider,
	albumRepository repositories.AlbumRepository,
	createMusicUsecase usecases.CreateMusicUsecase,
) *ReleaseAlbumUsecase {
	return &ReleaseAlbumUsecase{
		fileRepository,
		uuidProvider,
		albumRepository,
		createMusicUsecase,
	}
}

func (uc ReleaseAlbumUsecase) Release(params usecases.ReleaseAlbumParams) (*entities.Album, error) {
	coverLocation, err := uc.fileRepository.Upload(repositories.UploadFileParams{
		File:   params.CoverImageFile,
		Public: true,
	})

	if err != nil {
		return nil, err
	}

	releaseTime := time.Now()

	album, err := entities.NewAlbum(entities.NewAlbumParams{
		ID:            uc.uuidProvider.Generate(),
		Name:          params.Name,
		ReleaseDate:   releaseTime,
		OwnerID:       params.OwnerID,
		CoverImageURL: coverLocation.URL,
	})

	if err != nil {
		return nil, err
	}

	err = uc.albumRepository.Insert(album)

	if err != nil {
		return nil, err
	}

	var musics []*entities.Music
	for _, musicParams := range params.Musics {
		music, err := uc.createMusicUsecase.Create(usecases.CreateMusicParams{
			Name:         musicParams.Name,
			OwnerID:      params.OwnerID,
			AlbumID:      album.ID,
			DurationInMs: musicParams.DurationInMs,
			ReleaseDate:  releaseTime,
			CoverImage:   musicParams.CoverImageFile,
			Song:         musicParams.SongFile,
		})

		if err != nil {
			// TODO: rollback on this error case
			return nil, err
		}

		musics = append(musics, music)
	}

	album.Musics = musics

	return album, nil
}
