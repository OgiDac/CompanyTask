package router

import (
	"time"

	"github.com/OgiDac/CompanyTask/api/controllers"
	"github.com/OgiDac/CompanyTask/repository"
	"github.com/OgiDac/CompanyTask/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouter(timeout time.Duration, db *gorm.DB, public *gin.RouterGroup, private *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	uc := &controllers.UserController{
		UserUseCase: usecase.NewUserUseCase(ur, timeout),
	}

	publicGroup := public.Group("/users")
	privateGroup := private.Group("/users")

	publicGroup.GET("/", uc.GetAllUsers)
	publicGroup.POST("/login", uc.Login)
	publicGroup.POST("/", uc.CreateUser)
	privateGroup.PUT("/", uc.UpdateUser)
	privateGroup.DELETE("/:id", uc.DeleteUser)

}
