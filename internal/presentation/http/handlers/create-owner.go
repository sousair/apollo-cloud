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
	CreateOwnerRequest struct {
		Name string `json:"name" validate:"required"`
	}

	CreateOwnerResponse struct {
		Owner entities.Owner `json:"owner"`
	}

	CreateOwnerHttpHandler struct {
		createOwnerUsecase usecases.CreateOwnerUsecase
		validator          *validator.Validate
	}
)

var _ httputils.EchoHttpHandler = (*CreateOwnerHttpHandler)(nil)

func NewCreateOwnerHttpHandler(createOwnerUsecase usecases.CreateOwnerUsecase, validator *validator.Validate) *CreateOwnerHttpHandler {
	return &CreateOwnerHttpHandler{
		createOwnerUsecase: createOwnerUsecase,
		validator:          validator,
	}
}

func (h CreateOwnerHttpHandler) Handle(c echo.Context) error {
	var req CreateOwnerRequest

	if err := c.Bind(&req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return httputils.NewHttpErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	owner, err := h.createOwnerUsecase.Create(usecases.CreateOwnerParams{
		Name: req.Name,
	})

	if err != nil {
		// TODO: Improve error handling with custom errors from usecases and repositories.
		return httputils.NewHttpErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusCreated, CreateOwnerResponse{Owner: *owner})
}
