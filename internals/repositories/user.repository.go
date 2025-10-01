package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
) (string, error) {
	sql := `
		INSERT INTO
			follows(user_id, following_id)
		VALUES
			($1, $2)
	`

	ctag, err := ur.dbpool.Exec(ctx, sql, currUserId, targetUserid)
	if err != nil {
		return "", err
	}

	if ctag.Insert() {
		sql := `
			SELECT username
			FROM profiles
			WHERE user_id = $1
		`
		var uname string
		if err := ur.dbpool.QueryRow(ctx, sql, targetUserid).Scan(
			&uname,
		); err != nil {
			return "", err
		}

		return uname, nil
	}

	return "", errors.New("unable to find user")
}

func (ur *UserRepository) GetFollowingPost(
	ctx context.Context,
	currUserId int,
) ([]models.Post, error) {
	sql := `
		SELECT p.post_id, p.body, p.image, p.user_id, pr.username, p.created_at
		FROM follows f
		JOIN posts p ON
			p.user_id = f.following_id
		JOIN profiles pr ON
			f.following_id = p.user_id
		WHERE
			f.user_id = $1	
		ORDER BY
			p.created_at DESC
	`

	rows, err := ur.dbpool.Query(ctx, sql, currUserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		if err := rows.Scan(
			&post.PostID,
			&post.Content,
			&post.Image,
			&post.CreatorID,
			&post.CreatorUname,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (ur *UserRepository) CreatePostLike(
	ctx context.Context,
	currUid,
	postId int,
) error {
	sql := `
		INSERT INTO post_likes(user_id, post_id)
		VALUES ($1, $2)
	`
	ctag, err := ur.dbpool.Exec(ctx, sql, currUid, postId)
	if err != nil {
		return err
	}
	if !ctag.Insert() {
		return errors.New("pg unable to insert post like")
	}

	return nil
}

func (ur *UserRepository) CreatePostComment(
	ctx context.Context,
	currUid,
	postId int,
	commentBody string,
) error {
	sql := `
		INSERT INTO post_comments(user_id, post_id,comment)
		VALUES ($1, $2, $3)
	`

	ctag, err := ur.dbpool.Exec(ctx, sql, currUid, postId, commentBody)
	if err != nil {
		return err
	}
	if !ctag.Insert() {
		return errors.New("pg unable to insert post comment")
	}

	return nil
}
