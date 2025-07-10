package router

import (
	"time"

	"github.com/OgiDac/CompanyTask/api/controllers"
	"github.com/OgiDac/CompanyTask/config"
	"github.com/OgiDac/CompanyTask/repository"
	"github.com/OgiDac/CompanyTask/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewFileRouter(env *config.Env, timeout time.Duration, db *gorm.DB, mongoDB *mongo.Database, public *gin.RouterGroup) {
	// SQL User repo (to check user exists)
	userRepo := repository.NewUserRepository(db)

	// Mongo File repo
	fileRepo := repository.NewFileRepository(mongoDB)

	// Usecase with both
	fileUseCase := usecase.NewFileUseCase(userRepo, fileRepo, timeout)

	// Controller
	fileController := &controllers.FileController{
		FileUseCase: fileUseCase,
	}

	publicGroup := public.Group("/files")
	// Route
	publicGroup.POST("/:id/", fileController.UploadFile)
	publicGroup.GET("/:id/", fileController.DownloadFile)
	publicGroup.GET("/user/:id", fileController.GetFilesByUser)
	publicGroup.DELETE("/user/:id", fileController.DeleteFilesByUser)
}
