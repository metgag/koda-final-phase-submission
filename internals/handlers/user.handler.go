package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
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
