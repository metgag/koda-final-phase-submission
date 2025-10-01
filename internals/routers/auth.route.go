package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/final-assignment/internals/handlers"
	"github.com/metgag/final-assignment/internals/repositories"
)

func InitAuthRouter(r *gin.Engine, dbpool *pgxpool.Pool) {
	ar := repositories.NewAuthRepository(dbpool)
	ah := handlers.NewAuthHandler(ar)

	authGroup := r.Group("auth")
	{
		authGroup.POST("/register", ah.HandleRegister)
		authGroup.POST("/login", ah.HandleLogin)
	}
}
