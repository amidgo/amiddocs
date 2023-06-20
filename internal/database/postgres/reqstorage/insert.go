package reqstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/errorutils/reqerror"
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
	deleteRequestQuery = fmt.Sprintf(
		`
			DELETE FROM %s WHERE %s = $1
		`,
		reqmodel.RequestTable,
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

func (s *requestStorage) DeleteRequest(ctx context.Context, requestId uint64) error {
	cmdTag, err := s.p.Pool.Exec(ctx, deleteRequestQuery, requestId)
	if cmdTag.RowsAffected() == 0 {
		return reqerror.REQ_NOT_FOUND
	}
	if err != nil {
		return requestError(err, amiderrors.NewCause("delete from request by id", "DeleteRequest", _PROVIDER))
	}
	return nil
}
