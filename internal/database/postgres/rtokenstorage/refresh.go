package rtokenstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/internal/models/rtokenmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/google/uuid"
)

var (
	insertTokenQuery = fmt.Sprintf(
		`INSERT INTO %s (%s, %s, %s) VALUES ($1,$2,$3)`,
		rtokenmodel.RefreshTokenTable,

		rtokenmodel.SQL.UserId,
		rtokenmodel.SQL.Token,
		rtokenmodel.SQL.Expired,
	)
	updateTokenQuery = fmt.Sprintf(
		`UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3`,
		rtokenmodel.RefreshTokenTable,

		rtokenmodel.SQL.Token,
		rtokenmodel.SQL.Expired,

		rtokenmodel.SQL.UserId,
	)
	tokenByUserIdQuery = fmt.Sprintf(
		`SELECT %s, %s FROM %s WHERE %s = $1`,
		rtokenmodel.SQL.Token,
		rtokenmodel.SQL.Expired,

		rtokenmodel.RefreshTokenTable,

		rtokenmodel.SQL.UserId,
	)
	tokenQuery = fmt.Sprintf(
		`SELECT %s FROM %s WHERE %s = $1 AND %s = $2`,
		rtokenmodel.SQL.Expired,

		rtokenmodel.RefreshTokenTable,

		rtokenmodel.SQL.UserId,
		rtokenmodel.SQL.Token,
	)
)

// Upsert without expired check, if token not exist insert them else update
func (s *refreshTokenStorage) CreateAndSaveRefreshToken(ctx context.Context, userId uint64) (string, error) {
	_, err := s.TokenByUserId(ctx, userId)
	if amiderrors.Is(err, tokenerror.TOKEN_NOT_FOUND) {
		rtoken, err := s.InsertToken(ctx, userId)
		if err != nil {
			return "", rtokenError(err, amiderrors.NewCause("insert token", "CreateAndSaveRefreshToken", _PROVIDER))
		}
		return rtoken, nil
	}
	if err != nil {
		return "", rtokenError(err, amiderrors.NewCause("token by userId", "CreateAndSaveRefreshToken", _PROVIDER))
	}
	rtoken, err := s.UpdateTokenByUserId(ctx, userId)
	if err != nil {
		return "", rtokenError(err, amiderrors.NewCause("update token by userid", "CreateAndSaveRefreshToken", _PROVIDER))
	}
	return rtoken, nil
}

func (s *refreshTokenStorage) TokenByUserId(ctx context.Context, userId uint64) (*rtokenmodel.RefreshToken, error) {
	token := new(rtokenmodel.RefreshToken)
	token.UserId = userId
	err := s.p.Pool.QueryRow(
		ctx,
		tokenByUserIdQuery,
		userId,
	).Scan(&token.Token, &token.Expired)
	if err != nil {
		return nil, rtokenError(err, amiderrors.NewCause("get token by user id", "TokenByUserId", _PROVIDER))
	}
	return token, nil
}

func (s *refreshTokenStorage) Token(ctx context.Context, userId uint64, oldRefreshToken string) (*rtokenmodel.RefreshToken, error) {
	token := new(rtokenmodel.RefreshToken)
	token.UserId = userId
	err := s.p.Pool.QueryRow(
		ctx,
		tokenQuery,
		userId, oldRefreshToken,
	).Scan(&token.Expired)
	if err != nil {
		return nil, rtokenError(err, amiderrors.NewCause("get token by user id", "TokenByUserId", _PROVIDER))
	}
	return token, nil
}

func (s *refreshTokenStorage) UpdateTokenByUserId(ctx context.Context, userId uint64) (string, error) {
	token := uuid.New()
	expired := time.Now().Add(s.rtoken_time)
	_, err := s.p.Pool.Exec(
		ctx,
		updateTokenQuery,
		token, expired, userId,
	)
	if err != nil {
		return "", rtokenError(err, amiderrors.NewCause("update token by userId", "UpdateTokenByUserId", _PROVIDER))
	}
	return token.String(), nil
}

func (s *refreshTokenStorage) InsertToken(ctx context.Context, userId uint64) (string, error) {
	token := uuid.New()
	expired := time.Now().Add(s.rtoken_time)
	_, err := s.p.Pool.Exec(
		ctx,
		insertTokenQuery,
		userId, token, expired,
	)
	if err != nil {
		return "", rtokenError(err, amiderrors.NewCause("insert token", "InsertToken", _PROVIDER))
	}
	return token.String(), nil
}
