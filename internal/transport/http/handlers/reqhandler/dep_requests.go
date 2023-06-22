package reqhandler

import (
	"net/http"
	"strconv"

	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// DepartmentRequests godoc
//
//	@Summary		DepartmentRequests
//	@Description	get all requests with send status by dep id from query requires secretary access
//	@Tags			requests
//	@Accept			json
//	@Produce		json
//	@Param			depId	query		uint64	true	"department id"
//	@Success		200		{array}		reqmodel.RequestViewDTO
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		401		{object}	amiderrors.ErrorResponse
//	@Failure		403		{object}	amiderrors.ErrorResponse
//	@Failure		404		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/requests/by-department-id [get]
func (h *requestHandler) DepartmentRequests(c *fiber.Ctx) error {
	depId, err := strconv.ParseUint(c.Query("depId"), 10, 64)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("parse depId query", "DepartmentRequests", _PROVIDER))
	}
	reqList, err := h.reqprov.RequestListByDepartmentId(c.UserContext(), depId, reqfields.SEND)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(reqList)
}

// HistoryDepartentRequests godoc
//
//	@Summary		HistoryDepartentRequests
//	@Description	get all requests with done status by dep id from query requires secretary access
//	@Tags			requests
//	@Accept			json
//	@Produce		json
//	@Param			depId	query		uint64	true	"department id"
//	@Success		200		{array}		reqmodel.RequestViewDTO
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		401		{object}	amiderrors.ErrorResponse
//	@Failure		403		{object}	amiderrors.ErrorResponse
//	@Failure		404		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/requests/history-by-department-id [get]
func (h *requestHandler) HistoryDepartmentRequest(c *fiber.Ctx) error {
	depId, err := strconv.ParseUint(c.Query("depId"), 10, 64)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("parse depId query", "DepartmentRequests", _PROVIDER))
	}
	reqList, err := h.reqprov.RequestListByDepartmentId(c.UserContext(), depId, reqfields.DONE)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(reqList)
}
