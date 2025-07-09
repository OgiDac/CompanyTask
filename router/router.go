package router

import (
	"time"

	"github.com/OgiDac/CompanyTask/api/middleware"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func Setup(timeout time.Duration, db *gorm.DB, rabbitChanel *amqp.Channel, r *gin.Engine) {
	public := r.Group("/public/api")
	private := r.Group("/private/api", middleware.JwtAuthMiddleware("secret"))

	NewUserRouter(timeout, db, rabbitChanel, public, private)

}
