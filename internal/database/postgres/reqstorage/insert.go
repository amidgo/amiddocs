package reqstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	insertRequestQuery = fmt.Sprintf(
		`
		INSERT INTO %s
		(%s,%s,%s,%s,%s,%s)
		VALUES(
			(SELECT %s FROM %s WHERE %s = $1),
			$2,$3,$4,$5,$6
		)
		RETURNING %s
		`,
		reqmodel.RequestTable,

		// insert columns
		reqmodel.SQL.StatusId,
		reqmodel.SQL.Count,
		reqmodel.SQL.Date,
		reqmodel.SQL.UserID,
		reqmodel.SQL.DepartmentID,
		reqmodel.SQL.DocumentTypeId,

		// select id from request status
		reqmodel.SQL_STATUS.ID,
		reqmodel.RequestStatusTable,
		reqmodel.SQL_STATUS.Status,

		// returning id
		reqmodel.SQL.ID,
	)
)

func (s *requestStorage) InsertRequest(ctx context.Context, request *reqmodel.RequestDTO) (*reqmodel.RequestDTO, error) {
	err := s.p.Pool.QueryRow(
		ctx,
		insertRequestQuery,
		request.Status,
		request.Count,
		request.Date,
		request.UserID,
		request.DepartmentID,
		request.DocumentType.ID,
	).Scan(&request.ID)
	if err != nil {
		return nil, requestError(err, amiderrors.NewCause("insert request query", "InsertRequest", _PROVIDER))
	}
	return request, nil
}
