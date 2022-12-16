package handlers

import (
	"github.com/aryasadeghy/go-mic/data"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (pr *Product) GetProducts(c *gin.Context) {
	pd := data.GetProducts()
	c.JSON(http.StatusOK, pd)
}
func (pr *Product) AddProduct(c *gin.Context) {
	// gin parse request body and put it in prd
	var prd *data.Product
	if err := c.ShouldBindJSON(&prd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.AddProduct(prd)
	c.JSON(http.StatusCreated, prd)
}

func (pr *Product) UpdateProduct(c *gin.Context) {
	// gin get id from url
	id, err := strconv.Atoi(c.Param("productId"))

	// if id is not valid or error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}
	var prd *data.Product
	if err := c.ShouldBind(&prd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.UpdateProduct(id, prd)
	if err == data.ErrProductNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, prd)
}
