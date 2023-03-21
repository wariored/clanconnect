package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Try to get the token from the "Authorization" header
		header := c.Request().Header.Get("Authorization")
		token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer"))

		// If there's no token in the header, try to get it from the query parameters
		if token == "" {
			token = c.QueryParam("token")
		}
		log.Println("TOKEN SECTION")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is missing"})
		}

		// Check if token is valid
		claims, err := verifyToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		// Set user ID in context
		userID := claims["user_id"].(string)
		userRole := claims["role"].(string)
		c.Set("userID", userID)
		c.Set("userRole", userRole)

		return next(c)
	}
}


func RoleAccessMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Get("userRole").(string)
			if userRole == "" || role == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Not authorized"})
			}
			if userRole != role {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Not authorized"})
			}
			return next(c)
		}
	}
}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return the secret key
		return []byte("secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func DisableCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000", "http://localhost:3000"},
    })
}
