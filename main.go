package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.StaticFS("/assets", http.Dir("./static/assets"))
	router.LoadHTMLGlob("./static/html/*")

	router.GET("/index", LandingHandler)
	router.GET("/memo", LandingHandler)
	router.GET("/dateselect", LandingHandler)
	router.GET("/final", LandingHandler)

	router.Run(":8080")
}

func LandingHandler(g *gin.Context) {
	//get random hash which IS NOT from DB
	//return static only
	d := NewDbHandler("localhost")
	e := &(Event{"2222", "fooo"})
	d.SetEvent(e)
	g.HTML(http.StatusOK, "landing.tmpl", gin.H{
		"hash": "213132313",
	})
}

func MemoGetHandler(g *gin.Context) {
	//get hash from static by form POST
	//render memo
	g.HTML(http.StatusOK, "memo.tmpl", gin.H{})
}

func DateSelectHandler(g *gin.Context) {
	//get hash + info from memo by form POST
	//render date select
	g.HTML(http.StatusOK, "selectdate.tmpl", gin.H{})
}

func FinishHandler(g *gin.Context) {
	//get hash from date select by form POST
	//render link with hash
	g.HTML(http.StatusOK, "final.tmpl", gin.H{})
}
