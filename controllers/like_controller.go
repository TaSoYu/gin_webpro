package controllers

import (
	"exchangeapp/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId + ":likes"

	if err := global.RedisDB.Incr(ctx, likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully liked article",
	})
}

func GetLikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId + ":likes"

	likeValue, err := global.RedisDB.Get(ctx, likeKey).Result()

	if err == redis.Nil {
		likeValue = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"likes": likeValue,
	})
}
