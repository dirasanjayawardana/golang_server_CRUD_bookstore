// berisi file untuk inisialisai router

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	// buat masing-masing route
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Halo Dunia...") // mengirimkan response dalam bentuk string
	})

	// route auth
	InitAuthRouter(router, db)
	InitBookRouter(router, db)

	return router
}