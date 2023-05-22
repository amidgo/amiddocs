package departmenthandler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetAllWithTypes godoc
//
//	@Summary		GetAllDepartmentsWithTypes
//	@Description	get all departments with them doc types
//
//	@Tags			departments
//
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{array}		depmodel.DepartmentTypes
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Security		Token
//	@Router			/departments/get-all-types [get]
func (h *departmentHandler) GetAllWithTypes(c *fiber.Ctx) error {
	list, err := h.depP.DepartmentListWithTypes(c.UserContext())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(list)
}
