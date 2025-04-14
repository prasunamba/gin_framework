package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretkey = []byte("mysecretkey")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]string{
	//  convert it to database later
	"user1": "pass1",
	"user2": "pass2",
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {

	var user1cred Credentials

	if err := c.ShouldBindJSON(&user1cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pass := users[user1cred.Username]
	if pass != user1cred.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invaliduser",
		})
	}
	expirytimer := time.Now().Add(time.Minute * 5)
	claim := &Claims{
		Username: user1cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirytimer.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstring, err := token.SignedString(secretkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}
	c.SetCookie(
		"token",                                // name
		tokenstring,                            // value
		int(time.Until(expirytimer).Seconds()), // maxAge (seconds) — 1 hour
		"/",                                    // path
		"localhost",                            // domain
		false,                                  // secure
		true,                                   // httpOnly

	)
	c.JSON(http.StatusOK, gin.H{
		"msg": "coockie set",
	})

}
func GetHome(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "cookie not found",
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}
	claim := &Claims{}
	keyfunc := func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil

	}
	tkn, err := jwt.ParseWithClaims(token, claim, keyfunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": " token not extracted",
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid token",
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"msg": fmt.Sprintf("Hello %s", claim.Username),
	})

}
func Refreshtoken(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "cookie not found",
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}
	claim := &Claims{}
	keyfunc := func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil

	}
	tkn, err := jwt.ParseWithClaims(token, claim, keyfunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": " token not extracted",
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid token",
		})
	}

	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "request is not valid ",
		})
	}
	expirytimer := time.Now().Add(time.Minute * 5)
	claim.ExpiresAt = expirytimer.Unix()

	retoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstring, err := retoken.SignedString(secretkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}

	c.SetCookie(
		"refresh-token",                        // name
		tokenstring,                            // value
		int(time.Until(expirytimer).Seconds()), // maxAge (seconds) — 1 hour
		"/",                                    // path
		"localhost",                            // domain
		false,                                  // secure
		true,                                   // httpOnly

	)
	c.JSON(http.StatusOK, gin.H{
		"msg": "coockie got refreshed ",
	})
}
