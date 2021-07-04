package handler

import (
	"github.com/DenrianWeiss/catball/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AddRedirectRequest struct {
	Dest string `json:"dest"`
}

func Redirect(ctx *gin.Context) {
	path, _ := ctx.Params.Get("target")
	if path == "" {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	target := new(string)
	err := service.GetRedirect(path, target)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	ctx.Redirect(http.StatusFound, *target)
}

func ShowRedirect(ctx *gin.Context) {
	path, _ := ctx.Params.Get("target")
	if path == "" {
		ctx.String(http.StatusOK, "/")
		return
	}
	target := new(string)
	err := service.GetRedirect(path, target)
	if err != nil {
		ctx.String(http.StatusOK, "/")
		return
	}
	ctx.String(http.StatusOK, *target)
}

func AddRedirect(ctx *gin.Context) {
	token, _ := ctx.Params.Get("token")
	path, _ := ctx.Params.Get("path")
	to := &AddRedirectRequest{}
	err := ctx.ShouldBindJSON(to)
	if token != service.Config.AdminToken {
		ctx.String(http.StatusForbidden, "No token")
		return
	}
	if err != nil {
		ctx.String(http.StatusBadRequest, "Req denied")
		return
	}
	if strings.HasPrefix(to.Dest, "http") {
		service.AddRedirect(path, to.Dest)
		ctx.JSON(http.StatusOK, gin.H{
			"path": path,
			"dest": to.Dest,
		})
	}
}

func DelRedirect(ctx *gin.Context) {
	token, _ := ctx.Params.Get("token")
	path, _ := ctx.Params.Get("path")
	if token != service.Config.AdminToken {
		ctx.String(http.StatusForbidden, "No token")
		return
	}

	service.DelRedirect(path)
	ctx.String(http.StatusOK, "deleted %v", path)
}
