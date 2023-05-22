package departmenthandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// GetDepartmentById godoc
//
//	@Summary		Get Department By own ID
//	@Description	return department dto by id in param
//	@Tags			departments
//	@Accept			json
//	@Produce		json
//	@Param			id	query		uint64	true	"department id"
//	@Success		200	{object}	depmodel.DepartmentDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Security		Token
//	@Router			/departments/get-by-id [get]
func (h *departmentHandler) GetDepartmentById(c *fiber.Ctx) error {
	id, e := strconv.ParseUint(c.Query("id", "0"), 10, 64)
	if e != nil {
		return amiderrors.NewInternalErrorResponse(e, amiderrors.NewCause("parse id", "GetDepartmentById", _PROVIDER)).SendWithFiber(c)
	}

	dep, err := h.depP.DepartmentById(c.UserContext(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(dep)
}
