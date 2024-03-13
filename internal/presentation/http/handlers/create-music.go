package httphandlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
	httputils "github.com/sousair/apollo-cloud/internal/presentation/http/utils"
)

type (
	CreateMusicRequest struct {
		Name         string `form:"name" validate:"required"`
		OwnerID      string `form:"owner_id" validate:"required,uuid"`
		AlbumID      string `form:"album_id,omitempty" validate:"omitempty,uuid"`
		DurationInMs string `form:"duration_in_ms" validate:"required"`
		ReleaseDate  string `form:"release_date" validate:"required"`
	}

	CreateMusicResponse struct {
		Music *entities.Music `json:"music"`
	}

	CreateMusicHttpHandler struct {
		validator          *validator.Validate
		createMusicUsecase usecases.CreateMusicUsecase
	}
)

var _ httputils.EchoHttpHandler = (*CreateMusicHttpHandler)(nil)

func NewCreateMusicHttpHandler(validator *validator.Validate, createMusicUsecase usecases.CreateMusicUsecase) *CreateMusicHttpHandler {
	return &CreateMusicHttpHandler{
		validator,
		createMusicUsecase,
	}
}

func (h CreateMusicHttpHandler) Handle(c echo.Context) error {
	var req CreateMusicRequest

	if err := c.Bind(&req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	coverImageMultipartHeader, err := c.FormFile("cover_image")

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	coverImage, err := httputils.MultipartHeaderFileToOsFile(coverImageMultipartHeader)

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	defer coverImage.Close()

	musicMultipartHeader, err := c.FormFile("music_file")

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	musicFile, err := httputils.MultipartHeaderFileToOsFile(musicMultipartHeader)

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	defer musicFile.Close()

	durationInMs, err := strconv.Atoi(req.DurationInMs)
	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)

	music, err := h.createMusicUsecase.Create(usecases.CreateMusicParams{
		Name:         req.Name,
		OwnerID:      req.OwnerID,
		AlbumID:      req.AlbumID,
		DurationInMs: durationInMs,
		ReleaseDate:  releaseDate,
		CoverImage:   coverImage,
		MusicFile:    musicFile,
	})

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusCreated, CreateMusicResponse{Music: music})
}
