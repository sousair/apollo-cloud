package httphandlers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sousair/apollo-cloud/internal/domain/entities"
	"github.com/sousair/apollo-cloud/internal/domain/usecases"
	httputils "github.com/sousair/apollo-cloud/internal/presentation/http/utils"
)

type (
	GetMusicRequest struct {
		MusicID string `param:"id"`
		Owner   string `query:"owner" validate:"omitempty,oneof=true false"`
		Album   string `query:"album" validate:"omitempty,oneof=true false"`
	}

	GetMusicResponse struct {
		Music *entities.Music `json:"music"`
	}

	GetMusicHttpHandler struct {
		validator       *validator.Validate
		getMusicUsecase usecases.GetMusicUsecase
	}
)

var _ httputils.EchoHttpHandler = (*GetMusicHttpHandler)(nil)

func NewGetMusicHttpHandler(validator *validator.Validate, getMusicUsecase usecases.GetMusicUsecase) *GetMusicHttpHandler {
	return &GetMusicHttpHandler{
		validator,
		getMusicUsecase,
	}
}

func (h GetMusicHttpHandler) Handle(e echo.Context) error {
	var req GetMusicRequest

	if err := e.Bind(&req); err != nil {
		return httputils.NewHttpErrorResponse(e, http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return httputils.NewHttpErrorResponse(e, http.StatusBadRequest, "invalid request body")
	}

	includeOwnerData := false
	if req.Owner == "true" {
		includeOwnerData = true
	}

	includeAlbumData := false
	if req.Album == "true" {
		includeAlbumData = true
	}

	music, err := h.getMusicUsecase.Get(usecases.GetMusicParams{
		ID:               req.MusicID,
		IncludeAlbumData: includeAlbumData,
		IncludeOwnerData: includeOwnerData,
	})

	if err != nil {
		return httputils.NewHttpErrorResponse(e, http.StatusInternalServerError, "internal server error")
	}

	return e.JSON(http.StatusOK, GetMusicResponse{Music: music})
}
