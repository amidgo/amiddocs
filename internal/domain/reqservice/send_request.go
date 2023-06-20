package reqservice

import (
	"context"
	"time"

	"github.com/amidgo/amiddocs/internal/errorutils/doctypeerror"
	"github.com/amidgo/amiddocs/internal/errorutils/reqerror"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/amidtime"
)

func (s *requestService) SendRequest(
	ctx context.Context,
	roles []userfields.Role,
	req *reqmodel.CreateRequestDTO,
) (*reqmodel.RequestDTO, error) {

	err := s.checkRefreshTime(ctx, req.UserID, req.DocumentType)
	if err != nil {
		return nil, amiderrors.
			Wrap(err, amiderrors.NewCause("check refresh time", "SendRequest", _PROVIDER))
	}
	doctype, err := s.checkDocTypeRole(ctx, req.DocumentType, roles)
	if err != nil {
		return nil, amiderrors.
			Wrap(err, amiderrors.NewCause("check check doc type role", "SendRequest", _PROVIDER))
	}
	request := reqmodel.NewRequest(0, reqfields.SEND, req.Count, amidtime.Now(), req.UserID, req.DepartmentID, doctype)
	request, err = s.reqRepo.InsertRequest(ctx, request)
	if err != nil {
		return nil, amiderrors.
			Wrap(err, amiderrors.NewCause("insert request", "SendRequest", _PROVIDER))
	}
	return request, nil
}

func (s *requestService) checkDocTypeRole(ctx context.Context, dtype doctypefields.DocumentType, userRoles []userfields.Role) (*doctypemodel.DocumentTypeDTO, error) {
	docType, err := s.docTypeProv.DocTypeByType(ctx, dtype)
	if err != nil {
		return nil, err
	}
	for _, urole := range userRoles {
		for _, dtrole := range docType.Roles {
			if urole == dtrole {
				return docType, nil
			}
		}
	}
	return nil, doctypeerror.WRONG_USER_ROLE
}

func (s *requestService) checkRefreshTime(ctx context.Context, userId uint64, docType doctypefields.DocumentType) error {
	lastReq, err := s.reqProv.LastRequestByUserId(ctx, userId, docType)
	if amiderrors.Is(err, reqerror.REQ_NOT_FOUND) {
		return nil
	}
	if err != nil {
		return err
	}
	refreshDate := amidtime.Day * time.Duration(lastReq.DocumentType.RefreshTime)
	if lastReq.Date.T().Add(refreshDate).After(time.Now()) {
		return reqerror.REQ_REFRESH_DATE
	}
	return nil
}
