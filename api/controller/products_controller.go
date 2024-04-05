package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *ProductController) Create(c *gin.Context) {
	var products domain.Products

	err := c.ShouldBind(&products)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	products.ID = primitive.NewObjectID()

	products.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.ProductsUsecase.Create(c, &products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *ProductController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	products, err := u.ProductsUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
