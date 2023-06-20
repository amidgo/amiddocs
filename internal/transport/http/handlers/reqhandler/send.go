package reqhandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

// Send godoc
//
//	@Summary		Send request
//	@Description	"Send request to generate document"
//	@Tags			requests
//	@Accept			json
//	@Produce		json
//	@Param			request	body		reqmodel.CreateRequestDTO	true	"request"
//
//	@Success		201		{object}	reqmodel.RequestDTO
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		401		{object}	amiderrors.ErrorResponse
//	@Failure		403		{object}	amiderrors.ErrorResponse
//	@Failure		404		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/requests/send [post]
func (h *requestHandler) Send(c *fiber.Ctx) error {
	createReqDTO := new(reqmodel.CreateRequestDTO)

	err := c.BodyParser(createReqDTO)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("parse body", "Send", _PROVIDER))
	}

	err = validate.ValidateFields(createReqDTO)
	if err != nil {
		return err
	}

	id, err := h.jwtser.UserID(c)
	if err != nil {
		return err
	}
	createReqDTO.UserID = id
	roles, err := h.jwtser.UserRoles(c)
	if err != nil {
		return err
	}
	rq, err := h.reqser.SendRequest(c.UserContext(), roles, createReqDTO)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(rq)
}
