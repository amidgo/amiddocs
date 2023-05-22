package departmenthandler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetAllDepartments godoc
//
//	@Summary	Get All Departmnets
//	@Tags		departments
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		depmodel.DepartmentDTO
//	@Failure	400	{object}	amiderrors.ErrorResponse
//	@Failure	500	{object}	amiderrors.ErrorResponse
//	@Security	Token
//	@Router		/departments/get-all [get]
func (h *departmentHandler) GetAllDepartments(c *fiber.Ctx) error {
	departmentList, err := h.depP.DepartmentList(c.UserContext())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(departmentList)
}
