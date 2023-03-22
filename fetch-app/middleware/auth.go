package middleware

import (
	"encoding/json"
	"fetch-app/utils"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	Role  string
	Name  string
	Phone string
	jwt.RegisteredClaims
}

func AuthMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		utils.ResponseError(c, http.StatusUnauthorized, "no token provided")
		return
	}
	claims, err := parseToken(token, os.Getenv("SECRET"))
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if claims.Role == "" {
		utils.ResponseError(c, http.StatusForbidden, "forbidden")
		return
	}
	c.Set("claims", claims)
	c.Next()
}

func parseToken(tokenString, secret string) (claims AuthClaims, err error) {
	decodedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return AuthClaims{}, err
	}

	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid &&
		claims.VerifyExpiresAt(time.Now().Unix(), true) &&
		claims.VerifyIssuedAt(time.Now().Unix(), true) {

		authClaims := AuthClaims{}
		b, err := json.Marshal(claims)
		if err != nil {
			return AuthClaims{}, err
		}
		err = json.Unmarshal(b, &authClaims)
		if err != nil {
			return AuthClaims{}, err
		}
		return authClaims, nil
	}
	return AuthClaims{}, err
}

func OnlyAdmin(c *gin.Context) {
	claims, _ := c.Get("claims")
	if claims.(AuthClaims).Role != "admin" {
		utils.ResponseError(c, http.StatusForbidden, "forbidden")
		return
	}
	c.Next()
}
