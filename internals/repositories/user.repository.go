package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/final-assignment/internals/models"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
}

func NewUserRepository(dbpool *pgxpool.Pool) *UserRepository {
	return &UserRepository{dbpool: dbpool}
}

func (ur *UserRepository) CreateUserPost(
	ctx context.Context,
	postBody models.PostBody,
	imageName string,
	uid int,
) (string, error) {
	argIndex := 1
	var args []any
	var queryCols []string
	var queryVals []string

	queryCols = append(queryCols, "user_id")
	args = append(args, uid)

	queryVals = append(queryVals, fmt.Sprintf("$%d", argIndex))

	if postBody.Content != nil {
		queryCols = append(queryCols, "body")
		args = append(args, *postBody.Content)

		argIndex++
		queryVals = append(queryVals, fmt.Sprintf("$%d", argIndex))
	}
	if imageName != "" {
		queryCols = append(queryCols, "image")
		args = append(args, imageName)

		argIndex++
		queryVals = append(queryVals, fmt.Sprintf("$%d", argIndex))
	}

	sql := fmt.Sprintf(
		"INSERT INTO posts(%s) VALUES(%s)",
		strings.Join(queryCols, ","),
		strings.Join(queryVals, ","),
	)
	// log.Println(sql)

	ctag, err := ur.dbpool.Exec(ctx, sql, args...)
	if err != nil {
		return "", err
	}
	if !ctag.Insert() {
		return "", errors.New("pg insert post failed")
	}

	return "Create post succesfully", nil
}

func (ur *UserRepository) CreateUserFollowing(
	ctx context.Context,
	currUserId,
	targetUserid int,
) (pgconn.CommandTag, error) {
	sql := `
		INSERT INTO
			follows(user_id, following_id)
		VALUES
			($1, $2)
	`

	return ur.dbpool.Exec(ctx, sql, currUserId, targetUserid)
}
