package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/final-assignment/internals/handlers"
	"github.com/metgag/final-assignment/internals/middlewares"
	"github.com/metgag/final-assignment/internals/repositories"
	"github.com/redis/go-redis/v9"
)

func InitUserRouter(r *gin.Engine, dbpool *pgxpool.Pool, rdb *redis.Client) {
	ur := repositories.NewUserRepository(dbpool, rdb)
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
