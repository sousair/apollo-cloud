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
	GetOwnerAlbumsRequest struct {
		OwnerID string `param:"id" validate:"required,uuid"`
		Musics  string `query:"musics" validate:"omitempty,oneof=true false"`
	}

	GetOwnerAlbumsResponse struct {
		Albums []*entities.Album `json:"albums"`
	}

	GetOwnerAlbumsHandler struct {
		validator             *validator.Validate
		getOwnerAlbumsUsecase usecases.GetOwnerAlbumsUsecase
	}
)

var _ httputils.EchoHttpHandler = (*GetOwnerAlbumsHandler)(nil)

func NewGetOwnerAlbumsHandler(
	validator *validator.Validate,
	getOwnerAlbumsUsecase usecases.GetOwnerAlbumsUsecase,
) *GetOwnerAlbumsHandler {
	return &GetOwnerAlbumsHandler{
		validator,
		getOwnerAlbumsUsecase,
	}
}

func (h GetOwnerAlbumsHandler) Handle(c echo.Context) error {
	var req GetOwnerAlbumsRequest

	if err := c.Bind(&req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	includeMusicData := false
	if req.Musics == "true" {
		includeMusicData = true
	}

	albums, err := h.getOwnerAlbumsUsecase.GetAlbums(usecases.GetOwnerAlbumsParams{
		OwnerID:          req.OwnerID,
		IncludeMusicData: includeMusicData,
	})

	if err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, GetOwnerAlbumsResponse{Albums: albums})
}
