package controllers

import (
	"errors"
	"exchangeapp/global"
	"exchangeapp/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateExchangeRate(ctx *gin.Context) {
	var ExchangeRate models.ExchangeRate

	if err := ctx.ShouldBind(&ExchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ExchangeRate.Date = time.Now()
	fmt.Println("检验db  ", global.Db)
	if err := global.Db.AutoMigrate(&ExchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if err := global.Db.Create(&ExchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, ExchangeRate)
}

func GetExchangeRate(ctx *gin.Context) {
	var ExchangeRate []models.ExchangeRate

	if err := global.Db.Find(&ExchangeRate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
		}
		return
	}
	ctx.JSON(http.StatusOK, ExchangeRate)
}
