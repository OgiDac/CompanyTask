package router

import (
	"time"

	"github.com/OgiDac/CompanyTask/api/controllers"
	"github.com/OgiDac/CompanyTask/publisher"
	"github.com/OgiDac/CompanyTask/repository"
	"github.com/OgiDac/CompanyTask/usecase"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func NewUserRouter(timeout time.Duration, db *gorm.DB, rabbitChanel *amqp.Channel, public *gin.RouterGroup, private *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	publisher := publisher.NewRabbitPublisher(rabbitChanel, "user-queue")
	uc := &controllers.UserController{
		UserUseCase: usecase.NewUserUseCase(ur, publisher, timeout),
	}

	publicGroup := public.Group("/users")
	privateGroup := private.Group("/users")

	publicGroup.GET("/", uc.GetAllUsers)
	publicGroup.GET("/test-publish", uc.PublishTestEvent)
	publicGroup.POST("/login", uc.Login)
	publicGroup.POST("/", uc.CreateUser)
	privateGroup.PUT("/", uc.UpdateUser)
	privateGroup.DELETE("/:id", uc.DeleteUser)

}
