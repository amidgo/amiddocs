package reqservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *requestService) UpdateRequestStatus(ctx context.Context, reqId uint64, status reqfields.Status) error {
	req, err := s.reqProv.RequestById(ctx, reqId)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("find request by id", "UpdateRequestStatus", _PROVIDER))
	}
	if req.Status == status {
		return nil
	}
	err = s.reqRepo.UpdateRequestStatus(ctx, reqId, status)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("update request status", "UpdateRequestStatus", _PROVIDER))
	}
	return nil
}
