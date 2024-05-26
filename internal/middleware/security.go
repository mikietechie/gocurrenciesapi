package middleware

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
)

func GetUserFromC(c *gin.Context) *models.User {
	user := c.Value("user").(*models.User)
	return user
}

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		token := strings.Replace(bearerToken, "Bearer ", "", 1)
		verifiedToken, err := utils.VerifyToken(token)
		if err == nil {
			err = services.CheckBlackToken(*verifiedToken)
		}
		if err != nil {
			log.Println("WithAuth ", err)
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Set("token", *verifiedToken)
		c.Next()
	}
}

func WithUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token jwt.Token = c.Value("token").(jwt.Token)
		var user *models.User
		userID, err := token.Claims.GetSubject()
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		err = models.Db.Model(user).First(&user, userID).Error
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func AssertIsAdmin(c *gin.Context) error {
	if GetUserFromC(c).Role != "Admin" {
		return errors.New("only admins can access this section")
	}
	return nil
}

func WithAssertion(fn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := fn(c)
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func WithClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Query("apikey")
		var client *models.Client
		log.Println("WithClient Middleware:\tapiKey", apiKey)
		err := models.Db.Model(client).First(&client, "api_key = ?", apiKey).Error
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Set("client", client)
		if !client.HasReads() {
			c.Next()
		}
		go services.AddClientReadsUsed(client, 1)
	}
}
