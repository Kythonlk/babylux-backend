package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"go-backend/api/controller"
	"go-backend/bootstrap"
	"go-backend/domain"
	"go-backend/mongo"
	"go-backend/repository"
	"go-backend/usecase"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
