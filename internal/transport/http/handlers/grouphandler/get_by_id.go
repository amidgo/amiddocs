package grouphandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

const _ID_Q = "id"

// GetGroupById godoc
//
//	@Summary		Get Group By Id
//	@Description	Get GroupDTO by Id
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			id	query		uint64	true	"group id"
//	@Success		200	{object}	groupmodel.GroupDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Router			/groups/get-by-id [get]
func (h *GroupHandler) GetGroupById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Query(_ID_Q, "0"), 10, 64)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse group id", "GetGroupById", _PROVIDER)).SendWithFiber(c)
	}
	group, err := h.groupP.GroupById(c.UserContext(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(group)
}
