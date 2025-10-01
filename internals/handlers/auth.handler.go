package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/metgag/final-assignment/internals/models"
	"github.com/metgag/final-assignment/internals/pkg"
	"github.com/metgag/final-assignment/internals/repositories"
	"github.com/metgag/final-assignment/internals/utils"
)

type AuthHandler struct {
	ar *repositories.AuthRepository
}

func NewAuthHandler(ar *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{ar: ar}
}

func (ah *AuthHandler) HandleRegister(ctx *gin.Context) {
	var regBody models.RegisterBody

	if err := ctx.ShouldBindJSON(&regBody); err != nil {
		utils.LogCtxError(
			ctx, "UNABLE BIND REGISTER BODY", "Register body mismatch", err, http.StatusBadRequest,
		)
		return
	}

	hash := pkg.NewHashConfig()
	hash.UseRecommended()

	hashedPwd, err := hash.GenHash(regBody.Password)
	if err != nil {
		utils.LogCtxError(
			ctx, "SERVER UNABLE HASH REGISTER PASSWORD", "Internal server error", err, http.StatusInternalServerError,
		)
		return
	}

	regisUser, err := ah.ar.RegisterUser(ctx, regBody.Email, hashedPwd, regBody.Uname)
	if err != nil {
		utils.LogCtxError(
			ctx, "SERVER UNABLE REGISTER USER", "Email or username is taken", err, http.StatusInternalServerError,
		)
		return
	}

	ctx.JSON(http.StatusOK, models.NewFullfilledResponse(
		http.StatusOK,
		fmt.Sprintf(
			"Succesfully register %s", regisUser,
		),
	))
}

func (ah *AuthHandler) HandleLogin(ctx *gin.Context) {
	var loginBody models.LoginBody

	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
		utils.LogCtxError(
			ctx, "UNABLE BIND LOGIN BODY", "Login body mismatch", err, http.StatusBadRequest,
		)
		return
	}

	loginScan, err := ah.ar.LoginUser(ctx, loginBody.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			utils.LogCtxError(
				ctx, "NO MATCHING EMAIL", "Invalid email or password", err, http.StatusBadRequest,
			)
			return
		}
		utils.LogCtxError(
			ctx, "PG UNABLE TO GET EMAIL", "Internal server error", err, http.StatusInternalServerError,
		)
		return
	}

	hash := pkg.NewHashConfig()
	hash.UseRecommended()

	isMatch, err := hash.ComparePasswordAndHash(loginBody.Password, loginScan.HashPwd)
	if err != nil {
		utils.LogCtxError(
			ctx, "UNABLE TO HASH LOGIN PASSWORD", "Internal server error", err, http.StatusInternalServerError,
		)
		return
	}
	if !isMatch {
		utils.LogCtxError(
			ctx, "LOGIN PASSWORD MISMATCH", "Invalid email or password", errors.New("login password does not match"), http.StatusBadRequest,
		)
		return
	}

	claims := pkg.NewJWTClaims(loginScan.UID)
	token, err := claims.GenToken()
	if err != nil {
		utils.LogCtxError(
			ctx, "UNABLE GENERATE LOGIN ACCESS TOKEN", "Internal server error", err, http.StatusInternalServerError,
		)
		return
	}

	ctx.JSON(http.StatusOK, models.NewFullfilledResponse(
		http.StatusOK,
		models.LoginResponse{
			Message: fmt.Sprintf(
				"Logged in as %s", loginScan.Uname,
			),
			Token: token,
		},
	))
}
