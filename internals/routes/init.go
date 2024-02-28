// berisi file untuk inisialisai router

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// buat masing-masing route
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Halo Dunia...") // mengirimkan response dalam bentuk string
	})

	return router
}