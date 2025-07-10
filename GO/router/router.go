package router

import (
	"time"

	"github.com/OgiDac/CompanyTask/api/middleware"
	"github.com/OgiDac/CompanyTask/config"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func Setup(env *config.Env, timeout time.Duration, db *gorm.DB, rabbitChanel *amqp.Channel, r *gin.Engine) {
	public := r.Group("/public/api")
	private := r.Group("/private/api", middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	NewUserRouter(env, timeout, db, rabbitChanel, public, private)

}
