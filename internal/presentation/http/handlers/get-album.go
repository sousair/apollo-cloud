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
	GetAlbumRequest struct {
		AlbumID string `param:"id" validate:"required,uuid"`
		Musics  string `query:"musics" validate:"omitempty,oneof=true false"`
		Owner   string `query:"owner" validate:"omitempty,oneof=true false"`
	}

	GetAlbumResponse struct {
		Album *entities.Album `json:"album"`
	}

	GetAlbumHttpHandler struct {
		validator       *validator.Validate
		getAlbumUsecase usecases.GetAlbumUsecase
	}
)

var _ httputils.EchoHttpHandler = (*GetAlbumHttpHandler)(nil)

func NewGetAlbumHttpHandler(validator *validator.Validate, getAlbumUsecase usecases.GetAlbumUsecase) *GetAlbumHttpHandler {
	return &GetAlbumHttpHandler{
		validator,
		getAlbumUsecase,
	}
}

func (h GetAlbumHttpHandler) Handle(e echo.Context) error {
	var req GetAlbumRequest

	if err := e.Bind(&req); err != nil {
		return httputils.NewHttpErrorResponse(e, http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return httputils.NewHttpErrorResponse(e, http.StatusBadRequest, "invalid request body")
	}

	includeMusicsData := false
	if req.Musics == "true" {
		includeMusicsData = true
	}

	includeOwnerData := false
	if req.Owner == "true" {
		includeOwnerData = true
	}

	album, err := h.getAlbumUsecase.Get(usecases.GetAlbumParams{
		ID:                req.AlbumID,
		IncludeMusicsData: includeMusicsData,
		IncludeOwnerData:  includeOwnerData,
	})

	if err != nil {
		return httputils.NewHttpErrorResponse(e, http.StatusInternalServerError, "internal server error")
	}

	return e.JSON(http.StatusOK, GetAlbumResponse{Album: album})
}
