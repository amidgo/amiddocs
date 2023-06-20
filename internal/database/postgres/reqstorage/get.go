package reqstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
	"github.com/jackc/pgx/v5"
)

var selectMaxQuery = fmt.Sprintf(`
	SELECT 
		%s, %s, max(%s) 
	FROM 
		%s
	GROUP BY %s, %s`,
	// selectable values
	sqlutils.Full(reqmodel.SQL.UserID),
	sqlutils.Full(reqmodel.SQL.DocumentTypeId),
	// max date
	sqlutils.Full(reqmodel.SQL.Date),

	reqmodel.RequestTable,

	// group by
	sqlutils.Full(reqmodel.SQL.UserID),
	sqlutils.Full(reqmodel.SQL.DocumentTypeId),
)

// first arg doc type
// second arg user_id
var lastrequestQuery = fmt.Sprintf(
	`
	SELECT %s,%s,%s,%s,%s,%s,%s,%s,%s,array_agg(%s)
	FROM %s 
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = $1 
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s
		INNER JOIN (%s) AS max_date ON max_date.%s = $2 AND max_date.%s = %s
		GROUP BY %s, %s, %s
	`,
	// selectable variables
	sqlutils.Full(reqmodel.SQL.ID),
	sqlutils.Full(reqmodel.SQL_STATUS.Status),
	sqlutils.Full(reqmodel.SQL.Count),
	sqlutils.Full(reqmodel.SQL.Date),
	sqlutils.Full(reqmodel.SQL.UserID),
	sqlutils.Full(reqmodel.SQL.DepartmentID),
	sqlutils.Full(doctypemodel.SQL.ID),
	sqlutils.Full(doctypemodel.SQL.Type),
	sqlutils.Full(doctypemodel.SQL.RefreshTime),
	sqlutils.Full(usermodel.SQL_ROLES.Role),

	// from requests
	reqmodel.RequestTable,

	// inner join on doc type roles table by document_type_id
	doctypemodel.DocTypeRoleTable,
	sqlutils.Full(doctypemodel.SQL_ROLES.DocumentTypeId),
	sqlutils.Full(reqmodel.SQL.DocumentTypeId),

	// inner join on doc types by document type = $1
	doctypemodel.DocTypeTable,
	sqlutils.Full(doctypemodel.SQL.Type),

	// inner join roles
	usermodel.RolesTable,
	sqlutils.Full(usermodel.SQL_ROLES.ID),
	sqlutils.Full(doctypemodel.SQL_ROLES.RoleId),

	// inner join request status
	reqmodel.RequestStatusTable,
	sqlutils.Full(reqmodel.SQL.StatusId),
	sqlutils.Full(reqmodel.SQL_STATUS.ID),

	selectMaxQuery,
	reqmodel.SQL.UserID,
	reqmodel.SQL.DocumentTypeId,
	sqlutils.Full(doctypemodel.SQL.ID),

	// group by columns
	sqlutils.Full(reqmodel.SQL.ID),
	sqlutils.Full(reqmodel.SQL_STATUS.Status),
	sqlutils.Full(doctypemodel.SQL.ID),
)

func reqQuery(query string) string {
	return fmt.Sprintf(
		`
	SELECT %s,%s,%s,%s,%s,%s,%s,%s,%s,array_agg(%s)
	FROM %s 
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s 
		INNER JOIN %s ON %s = %s
		INNER JOIN %s ON %s = %s
	%s
		GROUP BY %s,%s, %s
		ORDER BY %s
	`,
		// selectable variables
		sqlutils.Full(reqmodel.SQL.ID),
		sqlutils.Full(reqmodel.SQL_STATUS.Status),
		sqlutils.Full(reqmodel.SQL.Count),
		sqlutils.Full(reqmodel.SQL.Date),
		sqlutils.Full(reqmodel.SQL.UserID),
		sqlutils.Full(reqmodel.SQL.DepartmentID),
		sqlutils.Full(doctypemodel.SQL.ID),
		sqlutils.Full(doctypemodel.SQL.Type),
		sqlutils.Full(doctypemodel.SQL.RefreshTime),
		sqlutils.Full(usermodel.SQL_ROLES.Role),

		// select from table
		reqmodel.RequestTable,

		// inner join on doctyperoles
		doctypemodel.DocTypeRoleTable,
		sqlutils.Full(doctypemodel.SQL_ROLES.DocumentTypeId),
		sqlutils.Full(reqmodel.SQL.DocumentTypeId),

		// inner join on doctypes
		doctypemodel.DocTypeTable,
		sqlutils.Full(reqmodel.SQL.DocumentTypeId),
		sqlutils.Full(doctypemodel.SQL.ID),

		// inner join roles
		usermodel.RolesTable,
		sqlutils.Full(usermodel.SQL_ROLES.ID),
		sqlutils.Full(doctypemodel.SQL_ROLES.RoleId),

		// inner join request status
		reqmodel.RequestStatusTable,
		sqlutils.Full(reqmodel.SQL.StatusId),
		sqlutils.Full(reqmodel.SQL_STATUS.ID),

		// query
		query,

		// group by columns
		sqlutils.Full(reqmodel.SQL.ID),
		sqlutils.Full(doctypemodel.SQL.ID),
		sqlutils.Full(reqmodel.SQL_STATUS.ID),

		// order by request date
		sqlutils.Full(reqmodel.SQL.Date),
	)
}

func scanRequest(row pgx.Row, req *reqmodel.RequestDTO) error {
	req.DocumentType = new(doctypemodel.DocumentTypeDTO)
	return row.Scan(
		&req.ID,
		&req.Status,
		&req.Count,
		&req.Date,
		&req.UserID,
		&req.DepartmentID,
		&req.DocumentType.ID,
		&req.DocumentType.Type,
		&req.DocumentType.RefreshTime,
		&req.DocumentType.Roles,
	)
}

func (s *requestStorage) LastRequestByUserId(
	ctx context.Context,
	userId uint64,
	docType doctypefields.DocumentType,
) (*reqmodel.RequestDTO, error) {
	req := new(reqmodel.RequestDTO)
	req.DocumentType = new(doctypemodel.DocumentTypeDTO)
	fmt.Println(lastrequestQuery)
	row := s.p.Pool.QueryRow(ctx, lastrequestQuery, docType, userId)
	err := scanRequest(row, req)
	if err != nil {
		return nil, requestError(err, amiderrors.NewCause("last req query scan", "LastRequestByUserId", _PROVIDER))
	}
	return req, nil
}

func (s *requestStorage) requestListByQuery(ctx context.Context, query string, args ...interface{}) ([]*reqmodel.RequestDTO, error) {
	rows, err := s.p.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, requestError(err, amiderrors.NewCause("get request rows query", "requestListByQuery", _PROVIDER))
	}
	defer rows.Close()
	reqList := make([]*reqmodel.RequestDTO, 0)
	for rows.Next() {
		req := new(reqmodel.RequestDTO)
		req.DocumentType = new(doctypemodel.DocumentTypeDTO)
		err = scanRequest(rows, req)
		if err != nil {
			return nil, requestError(err, amiderrors.NewCause("scan request from rows", "requestListByQuery", _PROVIDER))
		}
		reqList = append(reqList, req)
	}
	return reqList, nil
}

func (s *requestStorage) RequestListByDepartmentLimit(
	ctx context.Context,
	departmentId uint64,
	limit uint64,
	offset uint64,
) ([]*reqmodel.RequestDTO, error) {
	query := reqQuery("WHERE departments.id = $1") + "\nLIMIT = $2 OFFSET = $3"
	reqList, err := s.requestListByQuery(ctx, query, departmentId, limit, offset)
	if err != nil {
		return nil, amiderrors.Wrap(err, amiderrors.NewCause("get request list by query", "RequestListByDepartmentLimit", _PROVIDER))
	}
	return reqList, nil
}

func (s *requestStorage) RequestHistoryListByDepartmentIdLimit(
	ctx context.Context,
	departmentId uint64,
	limit uint64,
	offset uint64,
) ([]*reqmodel.RequestDTO, error) {
	query := reqQuery(
		fmt.Sprintf("WHERE %s = $1 AND %s = '%s'",
			sqlutils.Full(depmodel.SQL.ID),
			sqlutils.Full(reqmodel.SQL_STATUS.Status),
			reqfields.DONE),
	) + "\nLIMIT = $2 OFFSET = $3"
	reqList, err := s.requestListByQuery(ctx, query, departmentId, limit, offset)
	if err != nil {
		return nil, amiderrors.Wrap(err, amiderrors.NewCause("get request list by query", "RequestListByDepartmentIdHistoryLimit", _PROVIDER))
	}
	return reqList, nil
}

func (s *requestStorage) RequestById(ctx context.Context, reqId uint64) (*reqmodel.RequestDTO, error) {
	query := reqQuery(fmt.Sprintf("WHERE %s = $1", sqlutils.Full(reqmodel.SQL.ID)))
	req := new(reqmodel.RequestDTO)
	row := s.p.Pool.QueryRow(ctx, query, reqId)
	err := scanRequest(row, req)
	if err != nil {
		return nil, requestError(err, amiderrors.NewCause("get request query", "RequestById", _PROVIDER))
	}
	return req, nil
}

func (s *requestStorage) RequestListByUser(ctx context.Context, userId uint64) ([]*reqmodel.RequestDTO, error) {
	query := reqQuery(fmt.Sprintf("WHERE %s = $1", sqlutils.Full(reqmodel.SQL.UserID)))
	reqlist, err := s.requestListByQuery(ctx, query, userId)
	if err != nil {
		return nil, amiderrors.Wrap(err, amiderrors.NewCause("get request list by reqid", "RequestListByUser", _PROVIDER))
	}
	return reqlist, nil
}
