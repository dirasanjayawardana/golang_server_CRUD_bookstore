package routes

import (
	"golang_server_bookstore/internals/handlers"
	"golang_server_bookstore/internals/middlewares"
	"golang_server_bookstore/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	// buat group sub router
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	bookHandler:= handlers.InitBookHandler(bookRepo)

	// create book
	// localhost:4000/book/new
	bookRouter.POST("/new", middlewares.CheckToken ,bookHandler.CreateBook)

	// get books
	bookRouter.GET("", middlewares.CheckToken ,bookHandler.GetBooks)
}
