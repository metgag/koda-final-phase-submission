package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/metgag/final-assignment/internals/pkg"
	"github.com/metgag/final-assignment/internals/utils"
)

func ValidateToken(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		utils.MwareLogCtxError(
			ctx, "BEARER TOKEN NOT FOUND", "Need login", errors.New("no bearer token found"), http.StatusUnauthorized,
		)
		return
	}
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) != 2 {
		utils.MwareLogCtxError(ctx, "BEARER TOKEN UNRECOGNIZED", "Invalid bearer token", errors.New("bearer token unrecognized"), http.StatusUnauthorized)
		return
	}

	var claims pkg.Claims
	token := splitToken[1]
	if err := claims.VerifyToken(token); err != nil {
		if strings.Contains(err.Error(), jwt.ErrTokenInvalidIssuer.Error()) {
			utils.MwareLogCtxError(
				ctx, "JWT ISSUER UNRECOGNIZED", "Invalid bearer token", jwt.ErrTokenInvalidIssuer, http.StatusUnauthorized,
			)
			return
		}
		if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
			utils.MwareLogCtxError(
				ctx, "BEARER TOKEN EXPIRED", "Need login again", jwt.ErrTokenExpired, http.StatusUnauthorized,
			)
			return
		}
		utils.MwareLogCtxError(
			ctx, "VALIDATE TOKEN INTERNAL SERVER ERROR", "Internal server error", errors.New("validate token middleware error"), http.StatusInternalServerError,
		)
		return
	}

	ctx.Set("claims", claims)
	ctx.Next()
}
