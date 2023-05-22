package reqstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	updateRequestStatusQuery = fmt.Sprintf(
		`UPDATE %s SET %s = (SELECT %s FROM %s WHERE %s = $1) WHERE %s = $2 `,
		reqmodel.RequestTable,
		// update value
		reqmodel.SQL.StatusId,

		// select if from request status where status = $1
		reqmodel.SQL_STATUS.ID,
		reqmodel.RequestStatusTable,
		reqmodel.SQL_STATUS.Status,

		// where req id = $2
		reqmodel.SQL.ID,
	)
)

func (h *requestStorage) UpdateRequestStatus(ctx context.Context, reqId uint64, status reqfields.Status) error {
	_, err := h.p.Pool.Exec(ctx, updateRequestStatusQuery, status, reqId)
	if err != nil {
		return requestError(err, amiderrors.NewCause("update request query", "UpdateRequestStatus", _PROVIDER))
	}
	return nil
}
