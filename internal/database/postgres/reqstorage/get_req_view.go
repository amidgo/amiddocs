package reqstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
	"github.com/jackc/pgx/v5"
)

var (
	getRequestViewList = fmt.Sprintf(
		`
		SELECT %s, %s, %s, %s, %s, %s, %s
			FROM %s
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s

		WHERE %s = $1 AND %s = $2

		GROUP BY %s, %s,%s, %s

		ORDER BY %s
		`,
		// selectable variables
		sqlutils.Full(reqmodel.SQL.ID),
		sqlutils.Full(usermodel.SQL.Name),
		sqlutils.Full(usermodel.SQL.Surname),
		sqlutils.Full(usermodel.SQL.FatherName),
		sqlutils.Full(reqmodel.SQL_STATUS.Status),
		sqlutils.Full(doctypemodel.SQL.Type),
		sqlutils.Full(reqmodel.SQL.Count),

		// from request table
		reqmodel.RequestTable,

		// inner join users table
		usermodel.UserTable,
		sqlutils.Full(usermodel.SQL.ID),
		sqlutils.Full(reqmodel.SQL.UserID),

		// inner join document types table
		doctypemodel.DocTypeTable,
		sqlutils.Full(doctypemodel.SQL.ID),
		sqlutils.Full(reqmodel.SQL.DocumentTypeId),

		// inner join request status
		reqmodel.RequestStatusTable,
		sqlutils.Full(reqmodel.SQL_STATUS.ID),
		sqlutils.Full(reqmodel.SQL.StatusId),

		// department id
		sqlutils.Full(reqmodel.SQL.DepartmentID),
		sqlutils.Full(reqmodel.SQL_STATUS.Status),

		// group by
		sqlutils.Full(reqmodel.SQL.ID),
		sqlutils.Full(usermodel.SQL.ID),
		sqlutils.Full(doctypemodel.SQL.ID),
		sqlutils.Full(reqmodel.SQL_STATUS.Status),

		//order by
		sqlutils.Full(reqmodel.SQL.Date),
	)
)

func scanRequestViewDTO(row pgx.Row, req *reqmodel.RequestViewDTO) error {
	return row.Scan(
		&req.ID,
		&req.FIO.Name,
		&req.FIO.Surname,
		&req.FIO.FatherName,
		&req.Status,
		&req.DocumentType,
		&req.DocumentCount,
	)
}

func (s *requestStorage) RequestListByDepartmentId(ctx context.Context, depId uint64, status reqfields.Status) ([]*reqmodel.RequestViewDTO, error) {
	rows, err := s.p.Pool.Query(ctx, getRequestViewList, depId, status)
	if err != nil {
		return nil, requestError(err, amiderrors.NewCause("get all reqviewdto", "RequestListByDepartmentId", _PROVIDER))
	}
	reqList, err := sqlutils.ScanList(rows, scanRequestViewDTO)
	if err != nil {
		return nil, requestError(err, amiderrors.NewCause("get all scan request into list", "RequestListByDepartmentid", _PROVIDER))
	}
	return reqList, nil
}
