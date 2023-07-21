package clients

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEven(c *gin.Context) {
	auth := c.Query("auth")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	evenValue := GetValue(true)
	fmt.Println(evenValue, auth)
	c.JSON(http.StatusOK, evenValue)
}

func GetOdd(c *gin.Context) {
	auth := c.Query("auth")
	fmt.Println(auth)
	if auth == "" {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	oddValue := GetValue(false)
	fmt.Println(oddValue, auth)
	c.JSON(http.StatusOK, oddValue)
}

func GetValue(even bool) int {
	randomInt := rand.Intn(100)
	if even && randomInt%2 != 0 {
		return GetValue(even)
	} else if !even && randomInt%2 == 0 {
		return GetValue(even)
	}
	return randomInt
}
