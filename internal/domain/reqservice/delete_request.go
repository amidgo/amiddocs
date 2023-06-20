package reqservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
)

func (s *requestService) DeleteRequest(ctx context.Context, userId, requestId uint64) error {
	req, err := s.reqProv.RequestById(ctx, requestId)
	if err != nil {
		return err
	}
	if req.UserID != userId {
		return tokenerror.FORBIDDEN
	}
	if req.Status != reqfields.SEND {
		return nil
	}
	err = s.reqRepo.DeleteRequest(ctx, requestId)
	if err != nil {
		return err
	}
	return nil
}
