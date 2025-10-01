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
		post := userGroup.Group("posts")
		post.POST("", uh.HandleCreatePost)
		post.GET("", uh.HandleViewFollowedPost)
		post.POST("/like/:postId", uh.HandleLikePost)
		post.POST("/comment/:postId", uh.HandleCommentPost)

		userGroup.POST("/follow/:followId", uh.HandleFollowUser)
	}
}
