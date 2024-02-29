package handlers

import (
	"golang_server_bookstore/internals/models"
	"golang_server_bookstore/internals/repositories"
	"log"
	"net/http"
	"strconv"

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
		})
		return
	}

	// kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success get books",
		"data":     result,
	})
}

// ctx *gin.Context --> untuk mengambil request dan memberi response
func (item *BookHandler) GetBookById(ctx *gin.Context) {
	// ambil path variabel dengan nama id, dan konversi ke integer
	id, _ := strconv.Atoi(ctx.Param("id"))

	// cari buku berdasarkan id
	result, err := item.FindById(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// jika buku tidak ditemukan return not found
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}

	// kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success get book",
		"data":     result,
	})
}

// ctx *gin.Context --> untuk mengambil request dan memberi response
func (item *BookHandler) DeleteBookById(ctx *gin.Context) {
	// ambil path variabel dengan nama id, dan konversi ke integer
	id, _ := strconv.Atoi(ctx.Param("id"))

	// cari buku berdasarkan id
	result, err := item.FindById(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// jika buku tidak ditemukan return not found
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}

	// hapus buku berdasarkan id
	if err := item.DeleteById(id); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success delete book",
	})
}

// ctx *gin.Context --> untuk mengambil request dan memberi response
func (item BookHandler) CreateBook(ctx *gin.Context) {

	// buat struct body untuk menampung request dari body
	body := models.BookModel{}
	// ambil body,konversi dari json atau form ke struct
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
		})
		return
	}

	// kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success save book",
		"data":    result,
	})
}

// ctx *gin.Context --> untuk mengambil request dan memberi response
func (item *BookHandler) UpdateBookById(ctx *gin.Context) {
	// ambil path variabel dengan nama id, dan konversi ke integer
	id, _ := strconv.Atoi(ctx.Param("id"))

	// buat struct body untuk menampung request dari body
	body := models.BookModel{}
	// ambil body,konversi dari json atau form ke struct
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// cek apakah field pada body ada isinya atau tidak
	if body.Title == "" && (body.Description == nil || *body.Description != "") && body.Author == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "empty field, At least one field must be provided",
		})
		return
	}

	// cari buku berdasarkan id
	result, err := item.FindById(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// jika buku tidak ditemukan return not found
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}

	// update buku berdasarkan id
	if err := item.UpdateById(id, body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success update book",
	})
}
