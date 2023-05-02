package reqstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
)

func lastrequestQuery(query string) string {
	return fmt.Sprintf(
		`
	SELECT requests.id,request_status.status,requests.count,requests.date,requests.user_id,requests.department_id,
		document_types.id,document_types.type,document_types.refresh_time,roles.role
	FROM requests 
		INNER JOIN (select max(requests.date) AS max_date from requests) AS r ON requests.date = r.max_date
		INNER JOIN document_types ON requests.document_type_id = document_types.id 
		INNER JOIN roles ON roles.id = document_types.role_id
		INNER JOIN request_status ON requests.status_id = request_status.id
	%s
	`,
		query,
	)
}

func reqQuery(query string) string {
	return fmt.Sprintf(
		`
	SELECT requests.id,request_status.status,requests.count,requests.date,requests.user_id,requests.department_id,
		document_types.id,document_types.type,document_types.refresh_time,roles.role
	FROM requests 
		INNER JOIN document_types ON requests.document_type_id = document_types.id 
		INNER JOIN roles ON roles.id = document_types.role_id
		INNER JOIN request_status ON requests.status_id = request_status.id
	%s
		ORDER BY requests.date
	`,
		query,
	)
}

func scanRequest(row pgx.Row, req *reqmodel.RequestDTO) error {
	return row.Scan(
		&req.ID, &req.Status, &req.Count, &req.Date, &req.UserID, &req.DepartmentID,
		&req.DocumentType.ID, &req.DocumentType.Type, &req.DocumentType.RefreshTime, &req.DocumentType.Role,
	)
}

func (s *requestStorage) LastRequestByUserId(
	ctx context.Context,
	userId uint64,
	docType reqfields.DocumentType,
) (*reqmodel.RequestDTO, error) {
	req := new(reqmodel.RequestDTO)
	req.DocumentType = new(doctypemodel.DocumentTypeDTO)
	row := s.p.Pool.QueryRow(ctx, lastrequestQuery("WHERE requests.user_id = $1 AND document_types.type = $2"), userId, docType)
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

func (s *requestStorage) RequestListByDepartmentIdHistoryLimit(
	ctx context.Context,
	departmentId uint64,
	limit uint64,
	offset uint64,
) ([]*reqmodel.RequestDTO, error) {
	query := reqQuery("WHERE department.id = $1 AND request_status.status = 'DONE'") + "\nLIMIT = $2 OFFSET = $3"
	reqList, err := s.requestListByQuery(ctx, query, departmentId, limit, offset)
	if err != nil {
		return nil, amiderrors.Wrap(err, amiderrors.NewCause("get request list by query", "RequestListByDepartmentIdHistoryLimit", _PROVIDER))
	}
	return reqList, nil
}
