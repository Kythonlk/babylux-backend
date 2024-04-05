package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (tc *ProductController) Create(c *gin.Context) {
	var product domain.Product

	err := c.ShouldBind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	product.ID = primitive.NewObjectID()

	product.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.ProductUsecase.Create(c, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Product created successfully",
	})
}

func (u *ProductController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	products, err := u.ProductUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
