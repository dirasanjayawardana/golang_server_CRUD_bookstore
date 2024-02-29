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

	// localhost:4000/book
	// get books
	bookRouter.GET("", middlewares.CheckToken ,bookHandler.GetBooks)
	// get book by id
	bookRouter.GET("/:id", middlewares.CheckToken ,bookHandler.GetBookById)
	// delete book by id
	bookRouter.DELETE("/:id", middlewares.CheckToken ,bookHandler.DeleteBookById)
	// update book by id
	bookRouter.PATCH("/:id", middlewares.CheckToken ,bookHandler.UpdateBookById)
	// create book
	bookRouter.POST("/new", middlewares.CheckToken ,bookHandler.CreateBook)

}
