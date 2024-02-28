package handlers

import (
	"golang_server_bookstore/internals/models"
	"golang_server_bookstore/internals/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*repositories.BookRepo
}

func InitBookHandler(item *repositories.BookRepo) *BookHandler {
	return &BookHandler{item}
}

// ctx *gin.Context --> untuk mengambil request dan memberi response
func (item *BookHandler) GetBooks(ctx *gin.Context) {
	result, err := item.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		}) // kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
		return
	}

	// kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success get books",
		"data":     result,
	})
}

// ctx *gin.Context --> untuk mengambil request dan memberi response
func (item BookHandler) CreateBook(ctx *gin.Context) {
	// ambil body,konversi dari json atau form ke struct
	body := models.BookModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// simpan buku
	if err := item.SaveBook(body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ambil buku yang sudah tersimpan
	result, err := item.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		}) // kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
		return
	}

	// memberikan response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success save book",
		"data": result,
	})
}
