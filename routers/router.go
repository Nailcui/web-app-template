package routers

import (
	"app/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Eng *gin.Engine

func init() {
	logger.Info("router init")
	Eng = gin.Default()
	Eng.LoadHTMLGlob("static/template/*")
	Eng.StaticFS("/static/js", http.Dir("./static/js"))
	Eng.StaticFS("/static/css", http.Dir("./static/css"))
	Eng.Any("/", Index)
	Eng.Any("/index.html", Index)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "app index.html",
	})
}
