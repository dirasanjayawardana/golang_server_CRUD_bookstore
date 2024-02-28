package middlewares

import (
	"golang_server_bookstore/pkg"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ctx *gin.Context --> untuk mengambil request dan memberi response
func CheckToken(ctx *gin.Context) {

	// ambil header authorization
	bearerToken := ctx.GetHeader("Authorization")

	// cek apakah token ada
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Please login first",
		}) // akan menghentikan request, tidak melanjutkan ke handler
		return
	}

	// cek apakah token merupakan Bearer token
	if !strings.Contains(bearerToken, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Authorization",
		}) // akan menghentikan request, tidak melanjutkan ke handler
		return
	}

	// ambil token dari format Bearer token dan verifikasi token
	token := strings.Replace(bearerToken, "Bearer ", "", -1) // jika parameter terakhir kurang dari 1, maka akan mereplace semua
	_, err := pkg.VerifyToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "expierd") {
			log.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Session expired",
			})
		}
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// melanjutkan ke proses handler selanjutnya
	ctx.Next()
}
