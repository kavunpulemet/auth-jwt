package database

import (
	_ "embed"
	"errors"
	"github.com/jmoiron/sqlx"

	dbmodels "auth/internal/database/models"
	"auth/internal/utils"
	"database/sql"
)

type AuthorizationRepository interface {
	SaveRefreshToken(ctx utils.MyContext, refreshToken dbmodels.RefreshToken) error
	GetRefreshToken(ctx utils.MyContext, tokenID string) (string, error)
}

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

//go:embed sql/DeleteRefreshToken.sql
var deleteRefreshToken string

//go:embed sql/SaveRefreshToken.sql
var saveRefreshToken string

func (r *AuthorizationPostgres) SaveRefreshToken(ctx utils.MyContext, token dbmodels.RefreshToken) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.New("internal server error")
	}

	_, err = tx.ExecContext(ctx.Ctx, deleteRefreshToken, token.UserID)
	if err != nil {
		tx.Rollback()
		ctx.Logger.Error("failed to delete token: ", err)
		return errors.New("internal server error")
	}

	_, err = tx.ExecContext(ctx.Ctx, saveRefreshToken, token.ID, token.UserID, token.TokenHash, token.ExpiresAt)
	if err != nil {
		tx.Rollback()
		ctx.Logger.Error("failed to insert refresh token: ", err)
		return errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

//go:embed sql/GetRefreshToken.sql
var getRefreshToken string

func (r *AuthorizationPostgres) GetRefreshToken(ctx utils.MyContext, tokenID string) (string, error) {
	var tokenHash string

	err := r.db.GetContext(ctx.Ctx, &tokenHash, getRefreshToken, tokenID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("invalid access token")
		}

		ctx.Logger.Error("failed to get refresh token: ", err)
		return "", errors.New("internal server error")
	}

	return tokenHash, nil
}
