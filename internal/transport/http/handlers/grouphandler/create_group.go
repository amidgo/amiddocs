package grouphandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

// CreateGroup godoc
//
//	@Summary		CreateGroup
//	@Description	create group
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			groups	body		groupmodel.GroupDTO	true	"group dto"
//	@Success		200		{object}	groupmodel.GroupDTO
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		403		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/groups/create [post]
func (h *GroupHandler) CreateGroup(c *fiber.Ctx) error {
	group := new(groupmodel.GroupDTO)
	err := c.BodyParser(group)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse body", "CreateGroup", _PROVIDER))
	}

	err = validateGroup(group)
	if err != nil {
		return err
	}

	group, err = h.groupS.CreateGroup(c.UserContext(), group)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(group)
}

func validateGroup(group *groupmodel.GroupDTO) error {
	if group.EducationStartDate.T().After(group.EducationFinishDate.T()) {
		return grouperror.INVALID_EDUCATION_DATE
	}
	err := validate.ValidateFields(group)
	if err != nil {
		return err
	}
	return nil
}
