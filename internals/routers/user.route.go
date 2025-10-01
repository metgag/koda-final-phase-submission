package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/final-assignment/internals/handlers"
	"github.com/metgag/final-assignment/internals/middlewares"
	"github.com/metgag/final-assignment/internals/repositories"
)

func InitUserRouter(r *gin.Engine, dbpool *pgxpool.Pool) {
	ur := repositories.NewUserRepository(dbpool)
	uh := handlers.NewUserHandler(ur)

	userGroup := r.Group("user", middlewares.ValidateToken)
	{
		userGroup.POST("/post", uh.HandleCreatePost)
		userGroup.POST("/follow", uh.HandleFollowUser)
	}
}
