package middleware

import (
	"haikal/backend-api/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// ambil secret key dari .env
// jika tidak ada, gunakan default "secret_key"

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil header authorization dari request
		tokenString := c.GetHeader("Authorization")

		// jika token kosong, kembalikan respons 401 Unathorized
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort() // hentikan request selanjutnya
			return
		}

		// hapus prefix "Bearer " dari token
		// Header biasanya berbentuk: "Bearer <token>"
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// buat struct untuk menampung klaim token
		claims := &jwt.RegisteredClaims{}

		// parse token dan verifikasi tanda tangan dengan jwtKey
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// kembalikan kunci rahasia untuk memverifikasi token
			return jwtKey, nil
		})

		// jika token tidak valid atau terjadi error saat parsing
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "invalid token",
			})
			c.Abort() // hentikan request
			return
		}

		// simpat klaim "sub" (username) ke dalam context
		c.Set("username", claims.Subject)

		// lanjut ke handler berikutnya
		c.Next()
	}
}
