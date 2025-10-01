package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metgag/final-assignment/internals/models"
	"github.com/metgag/final-assignment/internals/pkg"
	"github.com/metgag/final-assignment/internals/repositories"
	"github.com/metgag/final-assignment/internals/utils"
)

type UserHandler struct {
	ur *repositories.UserRepository
}

func NewUserHandler(ur *repositories.UserRepository) *UserHandler {
	return &UserHandler{ur: ur}
}

func (uh *UserHandler) HandleCreatePost(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	user, _ := claims.(pkg.Claims)

	var postBody models.PostBody

	if err := ctx.ShouldBind(&postBody); err != nil {
		utils.LogCtxError(
			ctx, "UNABLE BIND POST BODY", "Post body mismatch", err, http.StatusBadRequest,
		)
		return
	}

	imagePost := postBody.Image
	var imageName string

	// handle if there is avail image input
	if imagePost != nil {
		ext := filepath.Ext(imagePost.Filename)
		filename := fmt.Sprintf("post_%d_%d%s",
			user.UserID, time.Now().Unix(), ext,
		)
		filepath := filepath.Join(
			"public", "post", filename,
		)

		if err := ctx.SaveUploadedFile(imagePost, filepath); err != nil {
			utils.LogCtxError(
				ctx, "INVALID POST IMAGE", "Unable handle image to post", err, http.StatusBadRequest,
			)
			return
		}
		imageName = filename
	}

	// abort if nothing to post
	if *postBody.Content == "" {
		utils.LogCtxError(ctx, "POST BODY EMPTY", "Input something to post", errors.New("unable create empty post"), http.StatusBadRequest)
		return
	}

	// execute insert post query
	result, err := uh.ur.CreateUserPost(
		ctx, postBody, imageName, user.UserID,
	)
	if err != nil {
		utils.LogCtxError(
			ctx, "UNABLE CREATE USER POST", "Internal server error", err, http.StatusInternalServerError,
		)
		return
	}

	ctx.JSON(http.StatusOK, models.NewFullfilledResponse(
		http.StatusCreated,
		result,
	))
}

func (uh *UserHandler) HandleFollowUser(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	user, _ := claims.(pkg.Claims)

	followingId, err := strconv.Atoi(ctx.Param("followId"))
	if err != nil {
		utils.LogCtxError(
			ctx, "POST ID SHOULD BE INTEGER", "Post not found", err, http.StatusBadRequest,
		)
		return
	}
      
	targetUname, err := uh.ur.CreateUserFollowing(
		ctx, user.UserID, followingId,
	)
	if err != nil {
		utils.LogCtxError(ctx, "UNABLE FOLLOW SAME USER", "User not found", err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, models.NewFullfilledResponse(
		http.StatusOK,
		fmt.Sprintf(
			"Succesfully following %s", targetUname,
		),
	))
}

func (uh *UserHandler) HandleViewFollowedPost(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	user, _ := claims.(pkg.Claims)

	posts, err := uh.ur.GetFollowingPost(ctx, user.UserID)
	if err != nil {
		utils.LogCtxError(ctx, "PG UNABLE GET POST LIST", "Internal server error", err, http.StatusInternalServerError)
		return
	}

	if len(posts) == 0 {
		utils.LogCtxError(
			ctx, "NOT FOLLOWING ANY OTHER USER", "Follow any user to view post", errors.New("user have no following"), http.StatusBadRequest,
		)
		return
	}

	ctx.JSON(http.StatusOK, models.NewFullfilledResponse(
		http.StatusOK,
		posts,
	))
}

func (uh *UserHandler) HandleLikePost(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	user, _ := claims.(pkg.Claims)

	postId, err := strconv.Atoi(ctx.Param("postId"))
	if err != nil {
		utils.LogCtxError(
			ctx, "POST ID SHOULD BE INTEGER", "Post not found", err, http.StatusBadRequest,
		)
		return
	}

	if err := uh.ur.CreatePostLike(ctx, user.UserID, postId); err != nil {
		utils.LogCtxError(ctx, "PG UNABLE CREATE POST LIKE", "Internal server error", err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, models.NewFullfilledResponse(
		http.StatusOK,
		fmt.Sprintf("Successfully like a post with ID, %d", postId),
	))
}

func (uh *UserHandler) HandleCommentPost(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	user, _ := claims.(pkg.Claims)

	postId, err := strconv.Atoi(ctx.Param("postId"))
	if err != nil {
		utils.LogCtxError(
			ctx, "POST ID SHOULD BE INTEGER", "Post not found", err, http.StatusBadRequest,
		)
		return
	}

	var commendBody models.CommentPostBody
	if err := ctx.ShouldBindJSON(&commendBody); err != nil {
		utils.LogCtxError(ctx, "COMMENT POST BODY MISMATCH", "Internal server error", err, http.StatusInternalServerError)
		return
	}

	if err := uh.ur.CreatePostComment(
		ctx, user.UserID, postId, commendBody.Comment,
	); err != nil {
		utils.LogCtxError(ctx, "PG ERROR CREATE POST COMMENT", "Post not found", err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, models.NewFullfilledResponse(
		http.StatusCreated,
		fmt.Sprintf(
			"Succesfully create comment to post id, %d", postId,
		),
	))
}
