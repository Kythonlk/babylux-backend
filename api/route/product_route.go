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

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewProductRepository(db, domain.CollectionProduct)
	tc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(tr, timeout),
	}
	group.GET("/products", tc.Fetch)
	group.POST("/product", tc.Create)
}
