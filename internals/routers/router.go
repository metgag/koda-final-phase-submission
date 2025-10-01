package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func InitRouter(p *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	r := gin.Default()

	InitAuthRouter(r, p)
	InitUserRouter(r, p, rdb)

	return r
}
