package departmenthandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

// CreateDepartment godoc
//
//	@Summary		Create department
//	@Description	enabled auto check unique name and shortName values
//	@Tags			departments
//	@Accept			json
//	@Produce		json
//	@Param			department	body		depmodel.DepartmentDTO	true	"department dto"
//	@Success		201			{object}	depmodel.DepartmentDTO
//	@Failure		401			{object}	amiderrors.ErrorResponse
//	@Failure		403			{object}	amiderrors.ErrorResponse
//	@Failure		500			{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Router			/departments/create [post]
func (h *DepartmentHandler) CreateDepartment(c *fiber.Ctx) error {

	dep := new(depmodel.DepartmentDTO)
	err := c.BodyParser(dep)
	if err != nil {
		amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse body", "CreateDepartment", _PROVIDER))
	}

	err = validate.ValidateFields(dep)
	if err != nil {
		return err
	}

	dep, err = h.depS.CreateDepartment(c.UserContext(), dep)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(dep)
}
