package handler

import (
	"fmt"
	"github.com/DenrianWeiss/catball/model"
	"github.com/DenrianWeiss/catball/service"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"html/template"
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
		err := service.AddRedirect(path, to.Dest)
		if err != nil {
			ctx.String(http.StatusBadRequest, "path exists.")
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"path": path,
				"dest": to.Dest,
			})
		}
	}
}

func DelRedirect(ctx *gin.Context) {
	token, _ := ctx.Params.Get("token")
	path, _ := ctx.Params.Get("path")
	if token != service.Config.AdminToken {
		ctx.String(http.StatusForbidden, "No token")
		return
	}

	err := service.DelRedirect(path)
	if err != nil {
		ctx.String(http.StatusForbidden, fmt.Sprintf("%s", err))
		return
	} else {
		ctx.String(http.StatusOK, "deleted %v", path)
	}
}

func AddDocument(ctx *gin.Context) {
	token, _ := ctx.Params.Get("token")
	path, _ := ctx.Params.Get("path")
	if token != service.Config.AdminToken {
		ctx.String(http.StatusForbidden, "No token")
		return
	}
	article := &model.Article{}
	err := ctx.ShouldBind(article)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "ill formed article")
		return
	}
	err = service.AddDocument(path, article)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "database error, %v", err)
		return
	}
	ctx.String(http.StatusOK, "added article %v", path)
}

func GetDocument(ctx *gin.Context) {
	path, _ := ctx.Params.Get("path")
	article := &model.Article{}
	err := service.GetDocument(path, article)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	md := markdown.ToHTML([]byte(article.Content), nil, nil)

	ctx.HTML(http.StatusOK, "document.html", gin.H{
		"title": article.Title,
		"body":  template.HTML(md),
	})

}

func RenderDocument(ctx *gin.Context) {
	path, _ := ctx.Params.Get("path")
	article := &model.Article{}
	err := service.GetDocument(path, article)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "Not found"
		})
		return
	}
	md := markdown.ToHTML([]byte(article.Content), nil, nil)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"title": article.Title,
		"author": article.Author,
		"body":  string(md),
	})
}
