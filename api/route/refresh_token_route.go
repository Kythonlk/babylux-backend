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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
