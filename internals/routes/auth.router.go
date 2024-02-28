package routes

import (
	"golang_server_bookstore/internals/handlers"
	"golang_server_bookstore/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {

	// buat group sub router
	authRouter := router.Group("/auth")
	authRepo := repositories.InitAuthRepo(db)
	authHandler := handlers.InitAuthHandler(authRepo)

	// register router
	// localhost:4000/auth/new
	authRouter.POST("/new", authHandler.Register)

	// login router
	authRouter.POST("", authHandler.Login)
}