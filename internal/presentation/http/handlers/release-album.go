package httphandlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
	httputils "github.com/sousair/apollo-cloud/internal/presentation/http/utils"
)

type (
	ReleaseAlbumRequest struct {
		Name    string `form:"name" validate:"required"`
		OwnerID string `form:"owner_id" validate:"required,uuid"`
	}

	ReleaseAlbumResponse struct {
		Album *entities.Album `json:"album"`
	}

	ReleaseAlbumHttpHandler struct {
		validator           *validator.Validate
		releaseAlbumUsecase usecases.ReleaseAlbumUsecase
	}
)

var _ httputils.EchoHttpHandler = (*ReleaseAlbumHttpHandler)(nil)

func NewReleaseAlbumHttpHandler(validator *validator.Validate, releaseAlbumUsecase usecases.ReleaseAlbumUsecase) *ReleaseAlbumHttpHandler {
	return &ReleaseAlbumHttpHandler{
		validator,
		releaseAlbumUsecase,
	}
}

func (h ReleaseAlbumHttpHandler) Handle(c echo.Context) error {
	var req ReleaseAlbumRequest

	if err := c.Bind(&req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	coverImage, err := c.FormFile("cover_image")

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	coverImageFile, err := httputils.MultipartHeaderFileToOsFile(coverImage)

	defer coverImageFile.Close()

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	musicsParams, err := h.getMusicData(c)

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	album, err := h.releaseAlbumUsecase.Release(usecases.ReleaseAlbumParams{
		Name:           req.Name,
		OwnerID:        req.OwnerID,
		CoverImageFile: coverImageFile,
		Musics:         musicsParams,
	})

	for _, musicParams := range musicsParams {
		if musicParams.CoverImageFile != nil {
			defer musicParams.CoverImageFile.Close()
		}

		defer musicParams.SongFile.Close()
	}

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusCreated, ReleaseAlbumResponse{Album: album})
}

func (h ReleaseAlbumHttpHandler) getMusicData(c echo.Context) ([]usecases.ReleaseAlbumMusicParams, error) {
	var musicsParams []usecases.ReleaseAlbumMusicParams

	seekMusic := true
	musicCounter := 1

	for seekMusic {
		musicName := c.FormValue("music_name_" + strconv.Itoa(musicCounter))

		if musicName == "" {
			seekMusic = false
			continue
		}

		durationInMs, err := strconv.Atoi(c.FormValue("music_duration_" + strconv.Itoa(musicCounter)))

		if err != nil {
			return nil, err
		}

		coverImage, err := c.FormFile("music_cover_" + strconv.Itoa(musicCounter))
		var coverImageFile *os.File

		if coverImage != nil {
			coverImageFile, err = httputils.MultipartHeaderFileToOsFile(coverImage)

			if err != nil {
				return nil, err
			}
		}

		songMultipartHeader, err := c.FormFile("music_song_" + strconv.Itoa(musicCounter))

		if err != nil {
			return nil, err
		}

		songFile, err := httputils.MultipartHeaderFileToOsFile(songMultipartHeader)

		if err != nil {
			return nil, err
		}

		musicsParams = append(musicsParams, usecases.ReleaseAlbumMusicParams{
			Name:           musicName,
			DurationInMs:   durationInMs,
			CoverImageFile: coverImageFile,
			SongFile:       songFile,
		})

		musicCounter++
	}

	return musicsParams, nil
}
