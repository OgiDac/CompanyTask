package router

import (
	"time"

	"github.com/OgiDac/CompanyTask/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(timeout time.Duration, db *gorm.DB, r *gin.Engine) {
	public := r.Group("/public/api")
	private := r.Group("/private/api", middleware.JwtAuthMiddleware("secret"))

	NewUserRouter(timeout, db, public, private)

}
