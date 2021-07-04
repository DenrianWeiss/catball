package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainPageHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
