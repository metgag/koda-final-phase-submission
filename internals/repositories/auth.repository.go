package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/final-assignment/internals/models"
)

type AuthRepository struct {
	dbpool *pgxpool.Pool
}

func NewAuthRepository(dbpool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{dbpool: dbpool}
}

func (ar *AuthRepository) RegisterUser(ctx context.Context, email, hashPwd, uname string) (string, error) {
	tx, err := ar.dbpool.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	sql := `
		INSERT INTO
			accounts (email, password)
		VALUES
			($1, $2)
		RETURNING user_id
	`

	var uid int
	if err := tx.QueryRow(ctx, sql, email, hashPwd).Scan(&uid); err != nil {
		return "", err
	}

	regisUname, err := ar.initUserProfile(tx, ctx, uid, uname)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return regisUname, nil
}

func (ar *AuthRepository) initUserProfile(tx pgx.Tx, ctx context.Context, uid int, uname string) (string, error) {
	sql := `
		INSERT INTO 
			profiles (user_id, username)
		VALUES
			($1, $2)
		RETURNING username
	`

	var regisUname string
	if err := tx.QueryRow(ctx, sql, uid, uname).Scan(&regisUname); err != nil {
		return "", err
	}

	return regisUname, nil
}

func (ar *AuthRepository) LoginUser(ctx context.Context, email string) (models.LoginScan, error) {
	sql := `
		SELECT a.user_id, a.password, p.username
		FROM accounts a
		JOIN profiles p ON 
			p.user_id = a.user_id
		WHERE email = $1
	`

	var loginScan models.LoginScan
	if err := ar.dbpool.QueryRow(ctx, sql, email).Scan(
		&loginScan.UID,
		&loginScan.HashPwd,
		&loginScan.Uname,
	); err != nil {
		return models.LoginScan{}, err
	}

	return loginScan, nil
}
