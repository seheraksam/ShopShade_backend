package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			jwtKey := os.Getenv("JWT_SECRET")
			fmt.Println("jwtKey on auth middleware", jwtKey)
			authHeader := c.GetHeader("Authorization")
			fmt.Println("authHeader on auth middleware", authHeader)
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
				c.Abort()
				return
			}

			tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			fmt.Println(tokenStr)
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtKey), nil
			})
			fmt.Println(token)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}
		claims := token.Claims.(jwt.MapClaims)	*/

		c.Set("user_id", "685d26286d83b9da387e291c")
		c.Set("role", "customer")
		c.Set("exp", "test_exp")
		c.Next()
	}
}
