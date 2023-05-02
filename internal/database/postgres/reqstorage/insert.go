package reqstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *requestStorage) InsertRequest(ctx context.Context, request *reqmodel.RequestDTO) (*reqmodel.RequestDTO, error) {
	err := s.p.Pool.QueryRow(
		ctx,
		`
		INSERT INTO requests
		(status_id,count,date,user_id,department_id,document_type_id)
		VALUES(
			(SELECT id FROM request_status WHERE status = $1),
			$2,$3,$4,$5,$6
		)
		RETURNING id
		`,
		request.Status, request.Count, request.Date, request.UserID, request.DepartmentID, request.DocumentType.ID,
	).Scan(&request.ID)
	if err != nil {
		return nil, requestError(err, amiderrors.NewCause("insert request query", "InsertRequest", _PROVIDER))
	}
	return request, nil
}
