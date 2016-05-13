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
	router.Run(":8080")
}

func LandingHandler(g *gin.Context) {
	//get random hash which IS NOT from DB
	//return static only
	d := NewDbHandler("localhost")
	e := &(Event{"2222", "fooo"})
	d.SetEvent(e)
	g.HTML(http.StatusOK, "index.tmpl", gin.H{
		"hash": "213132313",
	})
}
